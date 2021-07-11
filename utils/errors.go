package utils

import "fmt"

// InvalidPathError describe error if invalid path was provided as argument.
type InvalidPathError struct {
	Err  error
	Path string
}

// Error function provide error message.
func (e *InvalidPathError) Error() string {
	return fmt.Sprintf("invalid path (%s) provided: %s", e.Path, e.Err.Error())
}

// Unwrap function provide parent error.
func (e *InvalidPathError) Unwrap() error {
	return e.Err
}

// CannotScanDirError describe error which occur during parsing dir files.
type CannotScanDirError struct {
	Err error
}

// Error function provide error message.
func (e *CannotScanDirError) Error() string {
	return fmt.Sprintf("cannot scan directory: %s", e.Err.Error())
}

// Unwrap function provide parent error.
func (e *CannotScanDirError) Unwrap() error {
	return e.Err
}

// EnvDirIsEmptyError describe error if env dir has no files with environment variables.
type EnvDirIsEmptyError struct {
	Path string
}

// Error function provide error message.
func (e *EnvDirIsEmptyError) Error() string {
	return fmt.Sprintf("provided directory %s is empty", e.Path)
}

// Unwrap function provide parent error.
func (e *EnvDirIsEmptyError) Unwrap() error {
	return nil
}

// EnvDirIsNotExistsError describe error which occur if env dir is not exists in provided path.
type EnvDirIsNotExistsError struct {
	Path string
}

// Error function provide error message.
func (e *EnvDirIsNotExistsError) Error() string {
	return fmt.Sprintf("directory %s is not exists", e.Path)
}

// Unwrap function provide parent error.
func (e *EnvDirIsNotExistsError) Unwrap() error {
	return nil
}

// CannotOpenEnvDirError describe error which occur when cannot open env dir.
type CannotOpenEnvDirError struct {
	Err  error
	Path string
}

// Error function provide error message.
func (e *CannotOpenEnvDirError) Error() string {
	return fmt.Sprintf("cannot open file (%s): %s", e.Path, e.Err.Error())
}

// Unwrap function provide parent error.
func (e *CannotOpenEnvDirError) Unwrap() error {
	return e.Err
}

// CannotReadFileInfoError describe error which occur if cannot read file info for env dir.
type CannotReadFileInfoError struct {
	Err  error
	Path string
}

// Error function provide error message.
func (e *CannotReadFileInfoError) Error() string {
	return fmt.Sprintf("cannot read file info (%s): %s", e.Path, e.Err.Error())
}

// Unwrap function provide parent error.
func (e *CannotReadFileInfoError) Unwrap() error {
	return e.Err
}

// ProvidedPathIsNotDirError describe error which occur if provided path is not a directory.
type ProvidedPathIsNotDirError struct {
	Path string
}

// Error function provide error message.
func (e *ProvidedPathIsNotDirError) Error() string {
	return fmt.Sprintf("provided path %s is not a directory", e.Path)
}

// Unwrap function provide parent error.
func (e *ProvidedPathIsNotDirError) Unwrap() error {
	return nil
}
