package rerrors

// ErrMeta describes redeye-specific diagnostic information that is useful
// for inclusion in errors
type ErrMeta struct {
	// TODO
}

// IsMetadata satisfies the Metadata interface
// it is a dummy function that does nothing
func (ErrMeta) IsMetadata() {
	// does nothing
}
