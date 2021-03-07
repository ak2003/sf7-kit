#!/bin/sh
#
# This script will generate release log template based on the stg
# revision and prd revision for grab-pay service (services under grab-pay repo)
#
# How to use:
# just run this file with serviceName releaseTag rollbackTag
# for example:
# ./scripts/release-generate.sh paysi-history v2.0.131727.rekishi v2.0.130055.rekishi
# It will contains the changes of the service and grab-pay/common, grab-pay/external, grab-pay/public
#
# For long time no deploy service, you may need to exclude common
# then just add nocommon as the last param
# ./scripts/release-generate.sh paysi-history v2.0.131727.rekishi v2.0.130055.rekishi nocommon

if [[ "$#" -ne 1 && "$#" -ne 3 && "$#" -ne 4 ]]; then
  echo "invalid number of arguments, usage: ./scripts/release-generate.sh serviceName deploy_tag rollback_tag"
  exit 0
fi

V1=$2
V2=$3
ddLink=
gandalfLink=
echo ---
echo "= $1 ="
echo
echo "(NOTE) Release: [[ https://bitbucket.org/germentop/gt-kit/commits/tag/$V1 | $V1 ]]"
echo
echo "(WARNING) Rollback: [[ https://bitbucket.org/germentop/gt-kit/commits/tag/$V2 | $V2 ]]"
echo
if [[ $V1 = "" || $V2 = "" ]]; then
  echo "Error: failed to get git tags, if you are sure the tag is correct, it must due to you didn't rebase the master to get the latest tag"
  exit 1
fi
cd "$GOPATH"/src/gt-kit || exit
if [[ $(git tag -l $V1) == '' ]]; then
    echo "$V1"" isn't in your local branch, if you are sure the tag is correct, it must due to you didn't rebase the master to get the latest tag"
    exit
fi
echo Changes in project:
echo "| commit | author | description\\n| -----  | -----  | -----"
if [[ $4 == 'nocommon' ]]; then
    git --no-pager log "$V1"..."$V2" --pretty="| %h | //%an// | %s" --  "$1"
else
    git --no-pager log "$V1"..."$V2" --pretty="| %h | //%an// | %s" --  "$1" common external public
fi
echo
echo "Approver: "
case $1 in
    "paysi" )
    ddLink="https://app.datadoghq.com/dash/118754/paysi"
    gandalfLink="https://gandalf-ui.stg-myteksi.com/service-dashboard/paysiv1\\nhttps://gandalf-ui.stg-myteksi.com/service-dashboard/paysiv2"
    ;;
    "paysi-agent" )
    ddLink="https://app.datadoghq.com/dash/630256/paysi-agent"
    gandalfLink="https://gandalf-ui.stg-myteksi.com/service-dashboard/paysiagentint"
    ;;
    "paysi-history" )
    ddLink="https://app.datadoghq.com/dash/801853/paysi-history"
    gandalfLink="https://gandalf-ui.stg-myteksi.com/service-dashboard/paysihistory"
    ;;
    "paysi-notifier" )
    ddLink="https://app.datadoghq.com/dash/417313/paysi-notifier"
    gandalfLink="https://gandalf-ui.stg-myteksi.com/coverage-report notification service to partner, no endpoints"
    ;;
    "paysicore" )
    ddLink="https://app.datadoghq.com/dash/161472/paysicore"
    gandalfLink="https://gandalf-ui.stg-myteksi.com/service-dashboard/paysicore"
    ;;
    "grab-log" )
    ddLink="https://app.datadoghq.com/dash/246971/grablog"
    gandalfLink="https://gandalf-ui.stg-myteksi.com/coverage-report audit service, no logic endpoint"
    ;;
    "grab-money" )
    ddLink="https://app.datadoghq.com/dash/60153/grabmoney"
    gandalfLink="https://gandalf-ui.stg-myteksi.com/service-dashboard/grabmoney"
    ;;
    "paysicore-adapter" )
    ddLink="https://app.datadoghq.com/dash/677922/paysicore-adapter"
    gandalfLink="https://gandalf-ui.stg-myteksi.com/service-dashboard/paysicore-adapter"
    ;;
    "overlay" )
    ddLink="https://app.datadoghq.com/dash/284450/grabpay-overlay"
    gandalfLink="https://gandalf-ui.stg-myteksi.com/service-dashboard/overlay"
    ;;
    "paysi-grabcard" )
    ddLink="https://app.datadoghq.com/dashboard/t43-gxm-ann/paysi-grabcard"
    gandalfLink="https://gandalf-ui.stg-myteksi.com/service-dashboard/paysi-grabcard"
    ;;
esac
echo "Datadog link: "$ddLink
echo "Gandalf link: "$gandalfLink
echo
echo ---
