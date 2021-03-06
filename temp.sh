buildCheck=0
checkfiles=`find ./gobuild-*/build.check`
for eachfile in $checkfiles
do
    check=`cat $eachfile`
    echo "$eachfile : $check"
    if [ $check != "0" ]; then
      buildCheck=1
    fi   
done

echo $buildCheck > ./build.check

errfiles=`find ./gobuild-*/build.err`
for eachfile in $errfiles
do
    err=`cat $eachfile`
    echo "$eachfile : $err"
    echo $eachfile >> ./build.err
    echo "" >> ./build.err
    echo $err >> ./build.err
    echo "" >> ./build.err
done

A=$(awk -F/ '{print $2}' <<< './cb-artifact-gobuild-1.16/build.check')
echo $A

B=$(awk -F- '{print $4}' <<< $A)
echo $B

C="./cb-artifact-gobuild-1.16/build.err"
T="${C//build.err/build.check}"
echo $T