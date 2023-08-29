[@main]
您真的想去会议室?\
那里面可有很强大的怪物\
如果您确定要去,请点击进入\ \
<察看进入条件/@need>\
<进入会议室/@move>\
<取消/@exit>\


[@need]
如果您想进入会议室\
必须有100000黄金\ \
<取消/@exit>\


[@move]
#IF
checkgold 100000
#ACT
take 金币 100000
mapmove D516 87 15
break
#ELSESAY
对不起，您不够条件进入\ \
<取消/@exit>\