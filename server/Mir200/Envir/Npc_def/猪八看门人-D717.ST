[@main]
您真的想在下去一层?\
那里面可有很强大的怪物\
如果您确定要去,请点击进入\ \
<察看进入条件/@need>\
<进入八层/@move>\
<取消/@exit>\

[@need]
如果您想进入猪洞八层\
必须有50000黄金以证明您的实力\ \
<取消/@exit>\

[@move]
#IF
checkgold 50000
#ACT
take 金币 50000
mapmove D718 27 74
break
#ELSESAY
对不起，您不够条件进入\ \
<取消/@exit>\