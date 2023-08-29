[@main]
您真的想去地下矿区一层东?\
那里面可有很强大的怪物\
如果您确定要去,请点击进入\ \
<察看进入条件/@need>\
<进入东采矿场/@move>\
<取消/@exit>\


[@need]
如果您想进入地下矿区一层东\
必须有100000黄金\ \
<取消/@exit>\


[@move]
#IF
checkgold 100000
#ACT
take 金币 100000
mapmove D407 92 43
break
#ELSESAY
对不起，您不够条件进入\ \
<取消/@exit>\