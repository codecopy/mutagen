package daemon

import (
	"os"
	"os/signal"

	"github.com/pkg/errors"

	"github.com/spf13/cobra"

	"google.golang.org/grpc"

	"github.com/havoc-io/mutagen/cmd"
	"github.com/havoc-io/mutagen/pkg/daemon"
	"github.com/havoc-io/mutagen/pkg/grpcutil"
	"github.com/havoc-io/mutagen/pkg/ipc"
	daemonsvc "github.com/havoc-io/mutagen/pkg/service/daemon"
	promptsvc "github.com/havoc-io/mutagen/pkg/service/prompt"
	sessionsvc "github.com/havoc-io/mutagen/pkg/service/session"
	"github.com/havoc-io/mutagen/pkg/session"
)

func runMain(command *cobra.Command, arguments []string) error {
	// Validate arguments.
	if len(arguments) != 0 {
		return errors.New("unexpected arguments provided")
	}

	// Attempt to acquire the daemon lock and defer its release.
	lock, err := daemon.AcquireLock()
	if err != nil {
		return errors.Wrap(err, "unable to acquire daemon lock")
	}
	defer lock.Release()

	// Create a channel to track termination signals. We do this before creating
	// and starting other infrastructure so that we can ensure things terminate
	// smoothly, not mid-initialization.
	signalTermination := make(chan os.Signal, 1)
	signal.Notify(signalTermination, cmd.TerminationSignals...)

	// Create a session manager and defer its shutdown.
	sessionManager, err := session.NewManager()
	if err != nil {
		return errors.Wrap(err, "unable to create session manager")
	}
	defer sessionManager.Shutdown()

	// Create the gRPC server and defer its stoppage. We use a hard stop rather
	// than a graceful stop so that it doesn't hang on open requests.
	server := grpc.NewServer(
		grpc.MaxSendMsgSize(grpcutil.MaximumMessageSize),
		grpc.MaxRecvMsgSize(grpcutil.MaximumMessageSize),
	)
	defer server.Stop()

	// Create the daemon server, defer its shutdown, and register it.
	daemonServer := daemonsvc.NewServer()
	defer daemonServer.Shutdown()
	daemonsvc.RegisterDaemonServer(server, daemonServer)

	// Create and register the prompt server.
	promptsvc.RegisterPromptingServer(server, promptsvc.NewServer())

	// Create and register the session server.
	sessionsvc.RegisterSessionsServer(server, sessionsvc.NewServer(sessionManager))

	// Compute the path to the daemon IPC endpoint.
	endpoint, err := daemon.EndpointPath()
	if err != nil {
		return errors.Wrap(err, "unable to compute endpoint path")
	}

	// Create the daemon listener and defer its closure. Since we hold the
	// daemon lock, we preemptively remove any existing socket since it (should)
	// be stale.
	os.Remove(endpoint)
	listener, err := ipc.NewListener(endpoint)
	if err != nil {
		return errors.Wrap(err, "unable to create daemon listener")
	}
	defer listener.Close()

	// Serve incoming connections in a separate Goroutine, watching for serving
	// failure.
	serverErrors := make(chan error, 1)
	go func() {
		serverErrors <- server.Serve(listener)
	}()

	// Wait for termination from a signal, the daemon service, or the gRPC
	// server. We treat termination via the daemon service as a non-error.
	select {
	case sig := <-signalTermination:
		return errors.Errorf("terminated by signal: %s", sig)
	case <-daemonServer.Termination:
		return nil
	case err = <-serverErrors:
		return errors.Wrap(err, "daemon server termination")
	}
}

var runCommand = &cobra.Command{
	Use:          "run",
	Short:        "Run the Mutagen daemon",
	Hidden:       true,
	RunE:         runMain,
	SilenceUsage: true,
}

var runConfiguration struct {
	// help indicates whether or not help information should be shown for the
	// command.
	help bool
}

func init() {
	// Grab a handle for the command line flags.
	flags := runCommand.Flags()

	// Disable alphabetical sorting of flags in help output.
	flags.SortFlags = false

	// Manually add a help flag to override the default message. Cobra will
	// still implement its logic automatically.
	flags.BoolVarP(&runConfiguration.help, "help", "h", false, "Show help information")
}