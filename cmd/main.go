package main

import (
	mc "members_club"
	transport "members_club/transport"
)

func main() {
	member := make([]mc.Member, 0, 32)
	mc.Members = &member
	transport.Run("8000")
}
