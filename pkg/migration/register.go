package migration

// In this function we register the migrations for the Kluster CRD
// IMPORTANT: Don't remove migrations, don't reorder them!
// The position in the migrations slice is used for versioning,
// so the only thing that is sane is to append migrations
// to the end of the slice
func init() {
	defaultRegistry.migrations = []Migration{
		Init,
		AddAggregationLayerCertificates,
		// <-- Insert new migrations at the end only!
	}
}
