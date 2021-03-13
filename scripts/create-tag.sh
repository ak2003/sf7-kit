#!/bin/sh

DIFFID=$1
TAGID="order.diff.$DIFFID"
ORIGINURL="https://arikarniawan:Ak210185@bitbucket.org/germentop/gt-kit-stg.git"

echo "TAGID : $TAGID"
echo "DIFF : $DIFFID"

# crate tag in local before push to origin
if git rev-parse $TAGID >/dev/null 2>&1
then
    git tag -d $TAGID
fi
git tag $TAGID phabricator/diff/$DIFFID
git tag -d phabricator/diff/$DIFFID
git push $ORIGINURL $TAGID :phabricator/diff/$DIFFID
git push --delete $ORIGINURL phabricator/diff/$DIFFID