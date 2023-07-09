package restcli

import (
	"errors"

	"github.com/infinitybotlist/grevolt/rest"
	"github.com/infinitybotlist/grevolt/types"
)

// Fetch all server members.
//
// <target is the server id>
func (c *RestClient) FetchMembers(target string) (*types.MemberQueryResponse, error) {
	return rest.Request[types.MemberQueryResponse]{Path: "servers/" + target + "/members"}.With(&c.Config)
}

// Retrieve a member.
//
// <target is the server id>
// <member is the member id>
func (c *RestClient) FetchMember(target, member string) (*types.Member, error) {
	return rest.Request[types.Member]{Path: "servers/" + target + "/members/" + member}.With(&c.Config)
}

// Removes a member from the server.
//
// <target is the server id>
// <member is the member id>
func (c *RestClient) KickMember(target, member string) error {
	return rest.Request[types.APIError]{Method: rest.DELETE, Path: "servers/" + target + "/members/" + member}.NoContent(&c.Config)
}

// Edit a member by their id.
//
// <target is the server id>
// <member is the member id>
//
// <official docs use `server` instead of `target`, this is inconsistent and so grevolt uses target here>
func (c *RestClient) EditMember(target, member string, d *types.DataMemberEdit) (*types.Member, error) {
	return rest.Request[types.Member]{Method: rest.PATCH, Path: "servers/" + target + "/members/" + member, Json: d}.With(&c.Config)
}

// Query members by a given name, this API is not stable and will be removed in the future.
//
// # DEPRECATED
//
// <this will always return an error as it is deprecated *and* undocumented>
func (c *RestClient) QueryMembersByName(target string) error {
	return errors.New("this route is deprecated")
}

// Ban a user by their id.
//
// <target is the server id>
// <member is the member id>
//
// <official docs interchange `server` and `target`, this is inconsistent and so grevolt uses target/memebr here as commonly used elsewhere>
func (c *RestClient) BanUser(target, member string, d *types.DataBanCreate) (*types.ServerBan, error) {
	return rest.Request[types.ServerBan]{Method: rest.PUT, Path: "servers/" + target + "/bans/" + member, Json: d}.With(&c.Config)
}

// Remove a user's ban.
//
// <target is the server id>
// <member is the member id>
//
// <official docs interchange `server` and `target`, this is inconsistent and so grevolt uses target/memebr here as commonly used elsewhere>
func (c *RestClient) UnbanUser(target, member string) error {
	return rest.Request[types.APIError]{Method: rest.DELETE, Path: "servers/" + target + "/bans/" + member}.NoContent(&c.Config)
}

// Fetch all bans on a server.
//
// <target is the server id>
func (c *RestClient) FetchBans(target string) (*types.BanListResult, error) {
	return rest.Request[types.BanListResult]{Path: "servers/" + target + "/bans"}.With(&c.Config)
}

// Fetch all server invites.
//
// <target is the server id>
func (c *RestClient) FetchInvites(target string) (*types.InviteList, error) {
	return rest.Request[types.InviteList]{Path: "servers/" + target + "/invites"}.With(&c.Config)
}
