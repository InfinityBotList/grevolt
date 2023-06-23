#!/bin/bash

# Check for macOS, warn if not
if [ "$(uname)" != "Darwin" ]; then
  echo 'Warning: This script has only been tested on macOS.' >&2
fi

rm -rf tmp/go_client

# Look for swagger-codegen
if ! [ -x "$(command -v swagger-codegen)" ]; then
  echo 'Error: swagger-codegen is not installed.' >&2
  cat <<EOF
Linux users and users not using homebrew+macOS should consider making a wrapper script like below:

#!/bin/bash
export JAVA_HOME="${JAVA_HOME:-/opt/homebrew/opt/openjdk@11/libexec/openjdk.jdk/Contents/Home}"
exec "${JAVA_HOME}/bin/java"  -jar "/opt/homebrew/Cellar/swagger-codegen/3.0.46/libexec/swagger-codegen-cli.jar" "$@"
EOF
  exit 1
fi

# Remove struct
remove_struct() {
    sed -i '' "/type $1/,/}/d" types/types.go
}

replace_struct() {
    sed -i '' "s/*$1/*$2/g" types/types.go
}

# Remove and replace existing callers
remove_replace() {
    replace_struct $1 $2
    remove_struct $1
}

personalize() {
  sed -i '' "s/$1 \*Object/$1 \*$2/g" types/types.go
}

# Shorthand for remove_replace ABC File
is_file() {
    remove_replace $1 File
}

swagger-codegen generate \
   -i src-openapi/openapi.json \
   -l go \
   -o tmp/go_client \

cat tmp/go_client/model_* > types/types.go

cp types/types.go src-openapi/unpatched_types.patch

# Remove package fields, remove '' if on linux
sed -i '' '/package/d' types/types.go

# Remove /* ... */ comments, remove '' if on linux
sed -i '' '/\/\*/,/\*\//d' types/types.go

# Replace Permissions
sed -i '' 's/Permission Permission/PermissionFriendly PermissionFriendly/g' types/types.go
sed -i '' 's/type Permission string/type PermissionFriendly string/g' types/types.go

# Rename *AllOfBannedUserAvatar to *Attachment
remove_replace Attachment File # Replaced by File
is_file os.File
is_file Attachment
is_file AllOfMemberAvatar
is_file AllOfBannedUserAvatar
is_file AllOfUserAvatar
is_file AllOfServerIcon
is_file AllOfServerBanner
is_file AllOfUserProfileBackground
is_file AllOfWebhookAvatar

# Rename AllOfSnapshotWithContextContent to SnapshotContent
remove_replace AllOfSnapshotWithContextContent SnapshotContent

remove_struct SnapshotContent
remove_struct AllOfImageSize # WTF is this even supposed to be?
remove_struct AllOfMessageWebhook # We have MessageWebhook which is the same thing
remove_replace AllOfUserStatusPresence string
remove_struct AllOfUserStatus # UserStatus is way better than this
remove_struct AllOfUserProfile # UserProfile is way better than this
remove_struct Channel # Its empty, needs to be patched
remove_replace AllOfUserRelationship string
remove_replace AllOfDataEditUserStatus UserStatus
remove_replace AllOfDataEditServerSystemMessages ServerSystemMessages
remove_replace AllOfNewRoleResponseRole Role
remove_replace AllOfRevoltConfigFeatures RevoltFeatures
remove_replace AllOfMemberJoinedAt string
remove_replace AllOfMessageEdited string

# Replace SnapshotContent with *Object
replace_struct SnapshotContent Object

# Nit pick, but AllOf is annoying
sed -i '' 's/AllOf//g' types/types.go

# With AllOf out, we can make more tweaks and patches
remove_struct EmojiParent

# Another really annoying thing is *string, make this string
sed -i '' 's/\*string/string/g' types/types.go

# All int32's should be int64, swagger bug
sed -i '' 's/int32/int64/g' types/types.go

# Fix presence
sed -i '' 's/Presence string/Presence Presence/g' types/types.go
sed -i '' 's/type Presence Presence/type Presence string/g' types/types.go

# Fix SystemMessages *Object -> SystemMessages *ServerSystemMessages
personalize SystemMessages ServerSystemMessages

# More minor tweaks
personalize Avatar File
personalize Status UserStatus
personalize Icon File
personalize Banner File
personalize Profile UserProfile
personalize Bot BotInformation
personalize Relationship RelationshipStatus

# Finalize
cp src-openapi/parse.patch tmp/parse.patch
cat types/types.go >> tmp/parse.patch
mv tmp/parse.patch types/types.go

rm -rf tmp/go_client