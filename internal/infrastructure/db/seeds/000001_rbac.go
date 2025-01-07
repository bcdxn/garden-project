package seeds

import "github.com/segmentio/ksuid"

var (
	// role IDs
	guestRoleId = "rol_" + ksuid.New().String()
	userRoleId  = "rol_" + ksuid.New().String()
	vUserRoleId = "rol_" + ksuid.New().String()
	adminRoleId = "rol_" + ksuid.New().String()
	// action IDs
	readActionId   = "act_" + ksuid.New().String()
	listActionId   = "act_" + ksuid.New().String()
	createActionId = "act_" + ksuid.New().String()
	updateActionId = "act_" + ksuid.New().String()
	deleteActionId = "act_" + ksuid.New().String()
	// resource type IDs
	roleResId        = "res" + ksuid.New().String()
	actionResId      = "res" + ksuid.New().String()
	plantResId       = "res_" + ksuid.New().String()
	gardenResId      = "res_" + ksuid.New().String()
	userProfileResId = "res_" + ksuid.New().String()
)

var (
	Seed_000001_rbac = seed{
		Version: 1,
		Steps: []seedStep{
			{
				SQL: "INSERT INTO rbac_role(id, name) VALUES($1, $2)",
				Data: [][]any{
					{guestRoleId, "GUEST"},
					{userRoleId, "USER"},
					{vUserRoleId, "VERIFIED_USER"},
					{adminRoleId, "ADMIN"},
				},
			},
			{
				SQL: "INSERT INTO rbac_action(id, name) VALUES($1, $2)",
				Data: [][]any{
					{readActionId, "READ"},
					{listActionId, "LIST"},
					{createActionId, "CREATE"},
					{updateActionId, "UPDATE"},
					{deleteActionId, "DELETE"},
				},
			},
			{
				SQL: "INSERT INTO rbac_resource(id, name) VALUES($1, $2)",
				Data: [][]any{
					{plantResId, "PLANT"},
					{gardenResId, "GARDEN"},
					{userProfileResId, "USER_PROFILE"},
					{roleResId, "ROLE"},
					{actionResId, "ACTION"},
				},
			},
			{
				SQL: "INSERT INTO rbac_permission(role_id, action_id, resource_id) VALUES($1, $2, $3)",
				Data: [][]any{
					// unverified users
					{userRoleId, listActionId, plantResId},       // unverified users can list plant data
					{userRoleId, readActionId, plantResId},       // unverified users can read plant data
					{userRoleId, listActionId, gardenResId},      // unverified users can list garden data
					{userRoleId, readActionId, gardenResId},      // unverified users can read garden data
					{userRoleId, readActionId, userProfileResId}, // unverified users can read their user profile data
					// verified users
					{vUserRoleId, listActionId, plantResId},         // verified users can list plant data
					{vUserRoleId, readActionId, plantResId},         // verified users can read plant data
					{vUserRoleId, createActionId, plantResId},       // verified users can create plant data
					{vUserRoleId, updateActionId, plantResId},       // verified users can update plant data
					{vUserRoleId, deleteActionId, plantResId},       // verified users can delete plant data
					{vUserRoleId, listActionId, gardenResId},        // verified users can read garden data
					{vUserRoleId, readActionId, gardenResId},        // verified users can read garden data
					{vUserRoleId, createActionId, gardenResId},      // verified users can create garden data
					{vUserRoleId, updateActionId, gardenResId},      // verified users can update garden data
					{vUserRoleId, deleteActionId, gardenResId},      // verified users can delete garden data
					{vUserRoleId, readActionId, userProfileResId},   // verified users can read their user profile data
					{vUserRoleId, updateActionId, userProfileResId}, // verified users can update their user profile data
					{vUserRoleId, deleteActionId, userProfileResId}, // verified users can delete their user profile data
					// admin
					{adminRoleId, readActionId, plantResId},         // admin can read plant data
					{adminRoleId, listActionId, plantResId},         // admin can list plant data
					{adminRoleId, createActionId, plantResId},       // admin can create plant data
					{adminRoleId, updateActionId, plantResId},       // admin can update plant data
					{adminRoleId, deleteActionId, plantResId},       // admin can delete plant data
					{adminRoleId, readActionId, gardenResId},        // admin can read garden data
					{adminRoleId, listActionId, gardenResId},        // admin can read garden data
					{adminRoleId, createActionId, gardenResId},      // admin can create garden data
					{adminRoleId, updateActionId, gardenResId},      // admin can update garden data
					{adminRoleId, deleteActionId, gardenResId},      // admin can delete garden data
					{adminRoleId, readActionId, userProfileResId},   // admin can read user profile data
					{adminRoleId, listActionId, userProfileResId},   // admin can list user profile data
					{adminRoleId, createActionId, userProfileResId}, // admin can create user profile data
					{adminRoleId, updateActionId, userProfileResId}, // admin can update user profile data
					{adminRoleId, deleteActionId, userProfileResId}, // admin can delete user profile data
					{adminRoleId, readActionId, roleResId},          // admin can read rbac roles
					{adminRoleId, listActionId, roleResId},          // admin can list rbac roles
					{adminRoleId, createActionId, roleResId},        // admin can create rbac roles
					{adminRoleId, updateActionId, roleResId},        // admin can create rbac roles
					{adminRoleId, deleteActionId, roleResId},        // admin can create rbac roles
					{adminRoleId, readActionId, actionResId},        // admin can read rbac actions
					{adminRoleId, listActionId, actionResId},        // admin can list rbac actions
					{adminRoleId, createActionId, actionResId},      // admin can create rbac actions
					{adminRoleId, updateActionId, actionResId},      // admin can update rbac actions
					{adminRoleId, deleteActionId, actionResId},      // admin can delete rbac actions
				},
			},
		},
	}
)
