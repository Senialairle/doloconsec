import (
	"context"
	"fmt"
	"io"

	healthcare "google.golang.org/api/healthcare/v1"
)

// checkPreapprovals checks if the given dataset has preapprovals.
func checkPreapprovals(w io.Writer, projectID, location, datasetID string) error {
	ctx := context.Background()

	healthcareService, err := healthcare.NewService(ctx)
	if err != nil {
		return fmt.Errorf("healthcare.NewService: %v", err)
	}

	datasetsService := healthcareService.Projects.Locations.Datasets

	name := fmt.Sprintf("projects/%s/locations/%s/datasets/%s", projectID, location, datasetID)
	dataset, err := datasetsService.Get(name).Do()
	if err != nil {
		return fmt.Errorf("Get: %v", err)
	}

	if len(dataset.PreApprovedAccessRequests) > 0 {
		fmt.Fprintf(w, "Dataset %q has %d preapproved access requests.\n", dataset.Name, len(dataset.PreApprovedAccessRequests))
		return nil
	}
	fmt.Fprintln(w, "Dataset has no preapproved access requests.")
	return nil
}
  
