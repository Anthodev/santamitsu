# santamitsu
Golang Discord Bot to create a Secret Santa

## Commands

### For everyone
- `/info` - Get the information about the Secret Santa
- `/join` - Join the Secret Santa
- `/leave` - Leave the Secret Santa
- `/get-pair` - Get your match once the Secret Santa is drawn

### For designated moderators
- `/exclude member` - Exclude a member from the Secret Santa
- `/exclude pair` - Prevent a match between two members
- `/exclude remove` - Remove the exclusion for a member
- `/exclude remove-pair` - Remove the exclusion for a pair
- `/exclude list` - List the exclusions
- `/draw` - Draw and announce the pairs for the Secret Santa once it's locked

### For the server admin
- `/setup` - Setup the Secret Santa
- `/announce` - Announce the Secret Santa to the members on the servers
- `/cancel` - Cancel the Secret Santa
- `/delete` - Delete the Secret Santa
- `/lock` - Lock the participation for the Secret Santa
- `/unlock` - Unlock the participation for the Secret Santa
- `/moderator-role add-role` - Add a role to the list of moderator roles for the Secret Santa
- `/moderator-role remove-role` - Remove a role from the list of moderator roles for the Secret Santa
- `/moderator-role list-roles` - List the moderator roles for the Secret Santa