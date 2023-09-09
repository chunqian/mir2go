[@main]
#IF
checkpkpoint 2
#SAY
我将会把你传送到你想去的地方!\ \
<取消/@exit>\
#ELSEACT
goto @main-1

[@main-1]
欢迎使用我的传送服务.\
每传送1个地方将用掉你100黄金.\
你想去什么地方?\ \
<比奇城/@bwmove>  <泥城/@brdmove>  <比奇省/@ginmove>\
<毒蛇峡谷/@svmove>  <盟重/@mwmove>  <竞技场/@wwmove>\
<取消/@exit>\


[@bwmove]
#IF
checkgold 100
#ACT
take gold 100
mapmove 0 333 268
break
#ELSESAY
你没有足够的黄金使用我们的服务!\
快离开~不要浪费我的时间!\ \
<取消/@exit>\

[@brdmove]
#IF
checkgold 100
#ACT
take gold 100
mapmove 0 288 615
break
#ELSESAY
你没有足够的黄金使用我们的服务!\
快离开~不要浪费我的时间!\ \
<取消/@exit>\


[@ginmove]
#IF
checkgold 100
#ACT
take gold 100
mapmove 0 634 612
break
#ELSESAY
你没有足够的黄金使用我们的服务!\
快离开~不要浪费我的时间!\ \
<取消/@exit>\

[@svmove]
#IF
checkgold 100
#ACT
take gold 100
mapmove 2 504 484
break
#ELSESAY
你没有足够的黄金使用我们的服务!\
快离开~不要浪费我的时间!\ \
<取消/@exit>\

[@mwmove]
#IF
checkgold 100
#ACT
take gold 100
mapmove 3 330 330
break
#ELSESAY
你没有足够的黄金使用我们的服务!\
快离开~不要浪费我的时间!\ \
<取消/@exit>\

[@wwmove]
#IF
checkgold 100
#ACT
take gold 100
mapmove 11 177 324
break
#ELSESAY
你没有足够的黄金使用我们的服务!\
快离开~不要浪费我的时间!\ \
<取消/@exit>\