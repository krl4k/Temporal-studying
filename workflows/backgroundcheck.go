package workflows

import (
	"backgroundcheck/activities"
	"go.temporal.io/sdk/workflow"
	"time"
)

// BackgroundCheck is your custom Workflow Definition.
func BackgroundCheck(ctx workflow.Context, param string) (string, error) {
	// Define the Activity Execution options
	// StartToCloseTimeout or ScheduleToCloseTimeout must be set
	activityOptions := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, activityOptions)
	// Execute the Activity synchronously (wait for the result before proceeding)
	var ssnTraceResult string
	activityFuture := workflow.ExecuteActivity(ctx, activities.SSNTraceActivity, param)
	err := activityFuture.Get(ctx, &ssnTraceResult)
	if err != nil {
		return "", err
	}
	// Make the results of the Workflow available
	return ssnTraceResult, nil
}
