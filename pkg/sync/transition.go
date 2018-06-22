	"golang.org/x/text/unicode/norm"

	// the path given as the first argument with the digest specified by the
	// second argument.
	Provide(path string, digest []byte) (string, error)
}

type transitioner struct {
	root             string
	cache            *Cache
	symlinkMode      SymlinkMode
	recomposeUnicode bool
	provider         Provider
	problems         []*Problem
}

func (t *transitioner) recordProblem(path string, err error) {
	t.problems = append(t.problems, &Problem{Path: path, Error: err.Error()})
func (t *transitioner) ensureRouteWithProperCase(path string, skipLast bool) error {
	parent := t.root
		// Recompose content names if necessary.
		if t.recomposeUnicode {
			for i, name := range contents {
				contents[i] = norm.NFC.String(name)
			}
		}

func (t *transitioner) ensureExpectedFile(path string, expected *Entry) (os.FileMode, int, int, error) {
	cacheEntry, ok := t.cache.Entries[path]
		return 0, 0, 0, errors.New("unable to find cache information for path")
	info, err := os.Lstat(filepath.Join(t.root, path))
		return 0, 0, 0, errors.Wrap(err, "unable to grab file statistics")
		return 0, 0, 0, errors.Wrap(err, "unable to convert cached modification time format")
		return 0, 0, 0, errors.New("modification detected")
	}

	// Extract ownership.
	uid, gid, err := getOwnership(info)
	if err != nil {
		return 0, 0, 0, errors.Wrap(err, "unable to compute file ownership")
	return mode, uid, gid, nil
func (t *transitioner) ensureExpectedSymlink(path string, expected *Entry) error {
	target, err := os.Readlink(filepath.Join(t.root, path))
	if t.symlinkMode == SymlinkMode_SymlinkPortable {
func (t *transitioner) ensureNotExists(path string) error {
	_, err := os.Lstat(filepath.Join(t.root, path))
func (t *transitioner) removeFile(path string, target *Entry) error {
	if _, _, _, err := t.ensureExpectedFile(path, target); err != nil {
	return os.Remove(filepath.Join(t.root, path))
func (t *transitioner) removeSymlink(path string, target *Entry) error {
	if err := t.ensureExpectedSymlink(path, target); err != nil {
	return os.Remove(filepath.Join(t.root, path))
func (t *transitioner) removeDirectory(path string, target *Entry) bool {
	fullPath := filepath.Join(t.root, path)
		t.recordProblem(path, errors.Wrap(err, "unable to read directory contents"))
		return false
	}

	// Recompose content names if necessary.
	if t.recomposeUnicode {
		for i, name := range contentNames {
			contentNames[i] = norm.NFC.String(name)
		}
	unknownContentEncountered := false
			t.recordProblem(contentPath, errors.New("unknown content encountered on disk"))
			unknownContentEncountered = true
		// Handle content removal based on type.
			if !t.removeDirectory(contentPath, entry) {
				continue
			}
			if err = t.removeFile(contentPath, entry); err != nil {
				t.recordProblem(contentPath, errors.Wrap(err, "unable to remove file"))
				continue
			if err = t.removeSymlink(contentPath, entry); err != nil {
				t.recordProblem(contentPath, errors.Wrap(err, "unable to remove symlink"))
				continue
			t.recordProblem(contentPath, errors.New("unknown entry type found in removal target"))
			continue
		// At this point the removal must have succeeded, so remove the entry
		// from the target.
		delete(target.Contents, name)
	// If we didn't encounter any unknown content and the target contents are
	// empty, then we can attempt to remove the directory itself.
	if !unknownContentEncountered && len(target.Contents) == 0 {
			t.recordProblem(path, errors.Wrap(err, "unable to remove directory"))
			return false
	// Success.
	return true
func (t *transitioner) remove(path string, target *Entry) *Entry {
		return nil
	if err := t.ensureRouteWithProperCase(path, false); err != nil {
		t.recordProblem(path, errors.Wrap(err, "unable to verify path to target"))
		return target
	// Handle removal based on type.
		// Create a copy of target for mutation.
		targetCopy := target.Copy()

		// Attempt to reduce it.
		if !t.removeDirectory(path, targetCopy) {
			return targetCopy
		}
		if err := t.removeFile(path, target); err != nil {
			t.recordProblem(path, errors.Wrap(err, "unable to remove file"))
			return target
		if err := t.removeSymlink(path, target); err != nil {
			t.recordProblem(path, errors.Wrap(err, "unable to remove symlink"))
			return target
		t.recordProblem(path, errors.New("removal requested for unknown entry type"))
		return target
	return nil
func (t *transitioner) swapFile(path string, oldEntry, newEntry *Entry) error {
	fullPath := filepath.Join(t.root, path)
	if err := t.ensureRouteWithProperCase(path, false); err != nil {
	mode, uid, gid, err := t.ensureExpectedFile(path, oldEntry)
	// Compute the new file mode based on the new entry's executability.
	if newEntry.Executable {
		mode = MarkExecutableForReaders(mode)
	} else {
		mode = StripExecutableBits(mode)
	}
	// If both files have the same contents (differing only in permissions),
	// then we won't have staged the file, so we just change the permissions on
	// the existing file.
	if bytes.Equal(oldEntry.Digest, newEntry.Digest) {
	stagedPath, err := t.provider.Provide(path, newEntry.Digest)
	// Set the mode for the staged file.
	if err := os.Chmod(stagedPath, mode); err != nil {
		return errors.Wrap(err, "unable to set staged file mode")
	}

	// Set the ownership for the staged file.
	if err := setOwnership(stagedPath, uid, gid); err != nil {
		return errors.Wrap(err, "unable to set staged file ownership")
	}

func (t *transitioner) createFile(path string, target *Entry) error {
	// Ensure that the target path doesn't exist, e.g. due to a case conflict or
	// modification since the last scan.
	if err := t.ensureNotExists(path); err != nil {
		return errors.Wrap(err, "unable to ensure path does not exist")
	stagedPath, err := t.provider.Provide(path, target.Digest)
		return errors.Wrap(err, "unable to locate staged file")
	}

	// Compute the new file mode based on the new entry's executability.
	mode := newFileBaseMode
	if target.Executable {
		mode = MarkExecutableForReaders(mode)
	} else {
		mode = StripExecutableBits(mode)
	}

	// Set the mode for the staged file.
	if err := os.Chmod(stagedPath, mode); err != nil {
		return errors.Wrap(err, "unable to set staged file mode")
	if err := filesystem.RenameFileAtomic(stagedPath, filepath.Join(t.root, path)); err != nil {
		return errors.Wrap(err, "unable to relocate staged file")
	return nil
func (t *transitioner) createSymlink(path string, target *Entry) error {
	// Ensure that the target path doesn't exist, e.g. due to a case conflict or
	// modification since the last scan.
	if err := t.ensureNotExists(path); err != nil {
		return errors.Wrap(err, "unable to ensure path does not exist")
	if err := os.Symlink(target.Target, filepath.Join(t.root, path)); err != nil {
		return errors.Wrap(err, "unable to link to target")
	return nil
func (t *transitioner) createDirectory(path string, target *Entry) *Entry {
	// Ensure that the target path doesn't exist, e.g. due to a case conflict or
	// modification since the last scan.
	if err := t.ensureNotExists(path); err != nil {
		t.recordProblem(path, errors.Wrap(err, "unable to ensure path does not exist"))
		return nil
	if err := os.Mkdir(filepath.Join(t.root, path), newDirectoryBaseMode); err != nil {
		t.recordProblem(path, errors.Wrap(err, "unable to create directory"))
		return nil
			if c := t.createDirectory(contentPath, entry); c != nil {
				created.Contents[name] = c
			}
			if err := t.createFile(contentPath, entry); err != nil {
				t.recordProblem(contentPath, errors.Wrap(err, "unable to create file"))
			} else {
				created.Contents[name] = entry
			if err := t.createSymlink(contentPath, entry); err != nil {
				t.recordProblem(contentPath, errors.Wrap(err, "unable to create symlink"))
			} else {
				created.Contents[name] = entry
			t.recordProblem(contentPath, errors.New("creation requested for unknown entry type"))
	// Return the portion of the target that was created.
	return created
func (t *transitioner) create(path string, target *Entry) *Entry {
		return nil
		if err := os.MkdirAll(filepath.Dir(t.root), newDirectoryBaseMode); err != nil {
			t.recordProblem(path, errors.Wrap(err, "unable to create parent component of root path"))
			return nil
	if err := t.ensureRouteWithProperCase(path, true); err != nil {
		t.recordProblem(path, errors.Wrap(err, "unable to verify path to target"))
		return nil
	// Handle creation based on type.
		return t.createDirectory(path, target)
		if err := t.createFile(path, target); err != nil {
			t.recordProblem(path, errors.Wrap(err, "unable to create file"))
			return nil
			return target
		if err := t.createSymlink(path, target); err != nil {
			t.recordProblem(path, errors.Wrap(err, "unable to create symlink"))
			return nil
			return target
	} else {
		t.recordProblem(path, errors.New("creation requested for unknown entry type"))
		return nil
func Transition(
	root string,
	transitions []*Change,
	cache *Cache,
	symlinkMode SymlinkMode,
	recomposeUnicode bool,
	provider Provider,
) ([]*Entry, []*Problem) {
	// Create the transitioner.
	transitioner := &transitioner{
		root:             root,
		cache:            cache,
		symlinkMode:      symlinkMode,
		recomposeUnicode: recomposeUnicode,
		provider:         provider,
	}

			if err := transitioner.swapFile(t.Path, t.Old, t.New); err != nil {
				transitioner.recordProblem(t.Path, errors.Wrap(err, "unable to swap file"))
		// fails, record the reduced entry and continue to the next transition.
		if r := transitioner.remove(t.Path, t.Old); r != nil {
		// new entry is (or at least as much of it as we can create). If the new
		// entry is nil, this is a no-op.
		results = append(results, transitioner.create(t.Path, t.New))
	return results, transitioner.problems