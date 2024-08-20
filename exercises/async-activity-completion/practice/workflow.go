package async

import (
	"context"
	"encoding/hex"
	"time"

	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/workflow"
)

func Workflow(ctx workflow.Context, input string) (string, error) {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 300 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	logger := workflow.GetLogger(ctx)
	logger.Info("Async workflow started", "input", input)

	var result string

	err := workflow.ExecuteActivity(ctx, Activity, input).Get(ctx, &result)
	if err != nil {
		logger.Error("Activity failed.", "Error", err)
		return "", err
	}

	logger.Info("Async workflow completed.", "result", result)
	return result, nil
}

func Activity(ctx context.Context, input string) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("Activity", "input", input)
	// TODO Part A: Log the `taskToken` from your activity so you can use it
	// in another client. You may need to add an encoding library to your
	// list of imports so that your token can be logged to a terminal.
	activityInfo := activity.GetInfo(ctx)
	taskToken := activityInfo.TaskToken
	logger.Info("Activity", "taskToken", hex.EncodeToString(taskToken))

	// TODO Part B: Update this activity to return
	// activity.ErrResultPending so it can be completed asynchronously.
	return "Received " + input, activity.ErrResultPending
}
