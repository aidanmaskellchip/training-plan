package tools

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/pontus-devoteam/agent-sdk-go/pkg/tool"
	"strconv"
	"training-plan/internal/ai/client"
	"training-plan/internal/api/transport/request"
)

func toInt(v interface{}) int {
	if v == nil {
		return 0
	}
	switch v.(type) {
	case int:
		return v.(int)
	case float64:
		return int(v.(float64))
	case string:
		i, _ := strconv.Atoi(v.(string))
		return i
	default:
		return 0
	}
}

func NewCreateRunningProfileTool(apiClient *client.APIClient) *tool.FunctionTool {
	return tool.NewFunctionTool(
		"create_running_profile",
		"Create a new running profile for a user",
		func(ctx context.Context, params map[string]interface{}) (interface{}, error) {
			goalDistance, _ := params["goal_distance"].(string)
			terrain, _ := params["terrain"].(string)
			currentAbility, _ := params["current_ability"].(string)
			startDate, _ := params["start_date"].(string)
			goalDate, _ := params["goal_date"].(string)
			userID, _ := params["user_id"].(string)

			goalTime := toInt(params["goal_time"])
			current5k := toInt(params["current_5k"])
			current10k := toInt(params["current_10k"])
			runningDaysPerWeek := toInt(params["running_days_per_week"])
			longRunDay := toInt(params["long_run_day"])
			planLength := toInt(params["plan_length"])

			var currentHalfMarathon int
			if val, ok := params["current_half_marathon"]; ok {
				currentHalfMarathon = toInt(val)
			}

			var currentFullMarathon int
			if val, ok := params["current_full_marathon"]; ok {
				currentFullMarathon = toInt(val)
			}

			runningDaysParam, _ := params["running_days"].([]interface{})
			runningDays := make([]int, len(runningDaysParam))
			for i, v := range runningDaysParam {
				runningDays[i] = toInt(v)
			}

			userUUID, err := uuid.Parse(userID)
			if err != nil {
				return nil, fmt.Errorf("failed to parse user ID: %w", err)
			}

			req := request.CreateRunningProfileRequest{
				UserID:              userUUID,
				GoalDistance:        goalDistance,
				GoalTime:            goalTime,
				Terrain:             terrain,
				Current5K:           current5k,
				Current10K:          current10k,
				CurrentHalfMarathon: currentHalfMarathon,
				CurrentFullMarathon: currentFullMarathon,
				RunningDays:         runningDays,
				RunningDaysPerWeek:  runningDaysPerWeek,
				LongRunDay:          longRunDay,
				CurrentAbility:      currentAbility,
				PlanLength:          planLength,
				StartDate:           startDate,
				GoalDate:            goalDate,
			}

			err = apiClient.CreateRunningProfile(req)
			if err != nil {
				return nil, fmt.Errorf("failed to create running profile: %w", err)
			}

			return "Running profile created successfully.", nil
		},
	).WithSchema(map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"user_id": map[string]interface{}{
				"type":        "string",
				"description": "The ID of the user to create the running profile for",
			},
			"goal_distance": map[string]interface{}{
				"type":        "string",
				"description": "The goal distance for the running plan",
			},
			"goal_time": map[string]interface{}{
				"type":        "integer",
				"description": "The goal time for the running plan in minutes",
			},
			"terrain": map[string]interface{}{
				"type":        "string",
				"description": "The terrain for the running plan",
			},
			"current_5k": map[string]interface{}{
				"type":        "integer",
				"description": "The user's current 5k time in minutes",
			},
			"current_10k": map[string]interface{}{
				"type":        "integer",
				"description": "The user's current 10k time in minutes",
			},
			"current_half_marathon": map[string]interface{}{
				"type":        "integer",
				"description": "The user's current half marathon time in minutes",
			},
			"current_full_marathon": map[string]interface{}{
				"type":        "integer",
				"description": "The user's current full marathon time in minutes",
			},
			"running_days": map[string]interface{}{
				"type":        "array",
				"description": "The days the user is available to run",
				"items": map[string]interface{}{
					"type": "integer",
				},
			},
			"running_days_per_week": map[string]interface{}{
				"type":        "integer",
				"description": "The number of days per week the user wants to run",
			},
			"long_run_day": map[string]interface{}{
				"type":        "integer",
				"description": "The day of the week the user wants to do their long run",
			},
			"current_ability": map[string]interface{}{
				"type":        "string",
				"description": "The user's current running ability",
			},
			"plan_length": map[string]interface{}{
				"type":        "integer",
				"description": "The length of the running plan in weeks",
			},
			"start_date": map[string]interface{}{
				"type":        "string",
				"description": "The start date of the running plan",
			},
			"goal_date": map[string]interface{}{
				"type":        "string",
				"description": "The goal date of the running plan",
			},
		},
		"required": []string{"user_id", "goal_distance", "goal_time", "terrain", "current_5k", "current_10k", "running_days", "running_days_per_week", "long_run_day", "current_ability", "plan_length", "start_date", "goal_date"},
	})
}
