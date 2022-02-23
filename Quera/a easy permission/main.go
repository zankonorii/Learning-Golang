package main

type Permissions struct {
	canSeeMessages      bool
	canDeleteMessages   bool
	canEditMessages     bool
	canKickMembers      bool
	canMakeMembersAdmin bool
	canAddMembers       bool
}

func SetUserPermissions(permissions *Permissions) int8 {
	var data int8 = 0

	if permissions.canSeeMessages {
		data += 1
	}

	if permissions.canDeleteMessages {
		data += 2
	}

	if permissions.canEditMessages {
		data += 4
	}

	if permissions.canKickMembers {
		data += 8
	}

	if permissions.canMakeMembersAdmin {
		data += 16
	}

	if permissions.canAddMembers {
		data += 32
	}

	return int8(data);
}

func GetUserPermissions(permissionsMask int8) *Permissions {
	if permissionsMask > 63 || permissionsMask < 0 {
		return &Permissions{false, false, false, false, false, false}
	}

	var array[6]bool
	var i int = 0

	for permissionsMask > 0 {
		temp := permissionsMask % 2
		permissionsMask = permissionsMask / 2

		if temp == 1{
			array[i] = true
		}else{
			array[i] = false
		}
		i++
	}
	
	return &Permissions{array[0], array[1], array[2], array[3], array[4], array[5]}
}