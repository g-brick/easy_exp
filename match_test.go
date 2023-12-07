package simple_exp

import (
	"encoding/json"
	"testing"
)

type exampleRule struct {
	Keyword string `json:"keyword"`
}

var jsonExampleString = `[
    {
      "keyword": "P0||P1"
    },
    {
      "keyword": "P0||P1"
    },
    {
      "keyword": "ISIS邻居变化"
    },
    {
      "keyword": "BGP邻居中断"
    },
    {
      "keyword": "OSPF邻居建立"
    },
    {
      "keyword": "OSPF邻居中断"
    },
    {
      "keyword": "延时"
    },
    {
      "keyword": "\"错误码:26904\"||\"错误码:0\"&& ( UMON||Monitor||UMonNetwork||TianLiang )"
    },
    {
      "keyword": "UMON||UMonitor||UMonNetwork||TianLiang||UTopic-UMARKET"
    },
    {
      "keyword": "UDB"
    },
    {
      "keyword": "UDB"
    },
    {
      "keyword": "UDISK||TimeMachine-UHOST"
    },
    {
      "keyword": "UDISK||TimeMachine-UHOST"
    },
    {
      "keyword": "IPSecVPN-VPN||NAT网关||UNetFE-UNET||ULB||UFireWall||共享带宽"
    },
    {
      "keyword": "IPSecVPN-VPN||NAT网关||UNetFE-UNET||ULB||UFireWall||共享带宽"
    },
    {
      "keyword": "UVpcfe-UNET||UDPN||UVPCFE-UNET"
    },
    {
      "keyword": "UVpcfe-UNET||UDPN||UVPCFE-UNET"
    },
    {
      "keyword": "UBill-UBILL||UACCOUNT"
    },
    {
      "keyword": "UBill-UBILL||UACCOUNT"
    },
    {
      "keyword": "UMEM"
    },
    {
      "keyword": "UMEM"
    },
    {
      "keyword": "华为的板卡失效"
    },
    {
      "keyword": "\"错误码:130\"&& ( UHOST)"
    },
    {
      "keyword": "UHOST"
    },
    {
      "keyword": "UHADOOP||UDW||UDDP||UKafka"
    },
    {
      "keyword": "UHADOOP||UDW||UDDP||UKafka"
    },
    {
      "keyword": "USEC||UClean"
    },
    {
      "keyword": "USEC||UClean"
    },
    {
      "keyword": "板卡/设备故障"
    },
    {
      "keyword": "堆叠分裂||堆叠恢复"
    },
    {
      "keyword": "\"fail to find target ns\""
    },
    {
      "keyword": "电源失效"
    },
    {
      "keyword": "机房告警"
    },
    {
      "keyword": "设备硬件故障"
    },
    {
      "keyword": "ForwardEngine异常"
    },
    {
      "keyword": "路由数量限制"
    },
    {
      "keyword": "ARP表资源占用超75%"
    },
    {
      "keyword": "ACL资源不足"
    },
    {
      "keyword": "转发表资源占用率达到100%"
    },
    {
      "keyword": "DBA"
    },
    {
      "keyword": "P2&&!自建"
    },
    {
      "keyword": "UMON||UMonitor||UMonNetwork||TianLiang||UTopic-UMARKET"
    },
    {
      "keyword": "宕机迁移成功"
    },
    {
      "keyword": "ICP||备案系统"
    },
    {
      "keyword": "WBR端口DOWN"
    },
    {
      "keyword": "高危超量"
    },
    {
      "keyword": "UHostSec-USEC"
    },
    {
      "keyword": "UHostSec-USEC"
    },
    {
      "keyword": "UHOST"
    },
    {
      "keyword": "UFILE"
    },
    {
      "keyword": "UFILE"
    },
    {
      "keyword": "UES"
    },
    {
      "keyword": "UES"
    },
    {
      "keyword": "青岛"
    },
    {
      "keyword": "青岛"
    },
    {
      "keyword": "青岛"
    },
    {
      "keyword": "DBA"
    },
    {
      "keyword": "DEV"
    },
    {
      "keyword": "网络单点"
    },
    {
      "keyword": "UDDB"
    },
    {
      "keyword": "UDDB"
    },
    {
      "keyword": "DEV"
    },
    {
      "keyword": "DEV"
    },
    {
      "keyword": "DBA"
    },
    {
      "keyword": "DBA"
    },
    {
      "keyword": "WAF-WAF||USEC-WebScan||USSL"
    },
    {
      "keyword": "WAF-WAF||USEC-WebScan||USSL"
    },
    {
      "keyword": "mysql&&warning"
    },
    {
      "keyword": "mysql&&critical"
    },
    {
      "keyword": "mysql-master-ip-failover"
    },
    {
      "keyword": "process-monitor-mha"
    },
    {
      "keyword": "red"
    },
    {
      "keyword": "yellow||node"
    },
    {
      "keyword": "API-GW"
    },
    {
      "keyword": "延时"
    },
    {
      "keyword": "H3C/HW/RJ风扇失效"
    },
    {
      "keyword": "锐捷接口down"
    },
    {
      "keyword": "未报备"
    },
    {
      "keyword": "zookeeper"
    },
    {
      "keyword": "网易"
    },
    {
      "keyword": "InfluxdbGalaxy"
    },
    {
      "keyword": "InfluxDB"
    },
    {
      "keyword": "umonitor3_agentalive_umon"
    },
    {
      "keyword": "ISP封堵告警"
    },
    {
      "keyword": "elasticsearch&&critical"
    },
    {
      "keyword": "elasticsearch&&warning"
    },
    {
      "keyword": "USMS"
    },
    {
      "keyword": "【批次】"
    },
    {
      "keyword": "【巡检】"
    },
    {
      "keyword": "P2"
    },
    {
      "keyword": "IP地址冲突"
    },
    {
      "keyword": "ufile||UFILE||us3"
    },
    {
      "keyword": "ufile||UFILE||us3"
    },
    {
      "keyword": "uaek"
    },
    {
      "keyword": "uaek"
    },
    {
      "keyword": "uaccess||ucdn"
    },
    {
      "keyword": "uaccess||ucdn"
    },
    {
      "keyword": "umon"
    },
    {
      "keyword": "umon"
    },
    {
      "keyword": "152.32.193."
    },
    {
      "keyword": "152.32.193."
    },
    {
      "keyword": "udisk||udisk-host"
    },
    {
      "keyword": "数美"
    },
    {
      "keyword": "测试"
    },
    {
      "keyword": "UAPIGateway"
    },
    {
      "keyword": "UAPIGateway"
    },
    {
      "keyword": "DBA"
    },
    {
      "keyword": "DBA"
    },
    {
      "keyword": "create_fail"
    },
    {
      "keyword": "\"usns|usms|\"&&\"UMON_MPSM_usns_sms-container-memUsage\""
    },
    {
      "keyword": "DEV"
    },
    {
      "keyword": "DEV"
    },
    {
      "keyword": "DEV"
    },
    {
      "keyword": "\"usns|usms|umon-usms\""
    },
    {
      "keyword": "\"P2\"&&\"上海-自建\"||\"P2\"&&\"广州-自建\"||\"P2\"&&\"北京-自建\""
    },
    {
      "keyword": "\"description=sms-report\""
    },
    {
      "keyword": "【物理网-传输】"
    },
    {
      "keyword": "USms"
    },
    {
      "keyword": "USms"
    },
    {
      "keyword": "紫龙"
    },
    {
      "keyword": "\"fail to find target ns\""
    },
    {
      "keyword": "upq-usms"
    },
    {
      "keyword": "upq-usms"
    },
    {
      "keyword": "unoc&&warning"
    },
    {
      "keyword": "unoc&&critical"
    },
    {
      "keyword": "warning"
    },
    {
      "keyword": "critical"
    }
  ]`

func TestCompileMatchAllExample(t *testing.T) {
	var rules []*exampleRule
	_ = json.Unmarshal([]byte(jsonExampleString), &rules)
	for _, v := range rules {
		match, err := Compile(v.Keyword)
		if err != nil {
			t.Error(err)
			continue
		}
		target := "前  other_contents" + v.Keyword + " 后 another_contents  "
		t.Log("Matched by: ", target)
		t.Log("Matched res: ", match(target))
	}
}

func TestCompileMix1(t *testing.T) {
	expRule := "\"P2\"&&\"上海-自建\"||\"P2\"&&\"广州-自建\"||\"P2\"&&\"北京-自建\""
	t.Log("Expression rule: ", expRule)
	match, err := Compile(expRule)
	if err != nil {
		t.Error(err)
		return
	}
	target1 := "前  other_contents" + "广州-自建" + " 后 another_contents  " + "P2"
	t.Log("Matched by: ", target1)
	t.Log("Matched res: ", match(target1))
	target2 := "前  other_contents" + "广州-自建" + " 后 another_contents  " + "P1"
	t.Log("Matched by: ", target2)
	t.Log("Matched res: ", match(target2))
}

func TestCompileMix2(t *testing.T) {
	expRule := "\"错误码:26904\"||\"错误码:0\"&& ( UMON||Monitor||UMonNetwork||TianLiang )"
	t.Log("Expression rule: ", expRule)

	match, err := Compile(expRule)
	if err != nil {
		t.Error(err)
		return
	}
	target1 := "前  other_contents" + "错误码:26904" + " 后 another_contents  "
	t.Log("Matched by: ", target1)
	t.Log("Matched res: ", match(target1))
	target2 := "前  other_contents" + "错误码:0" + " 后 another_contents  " + "UMonNetwork"
	t.Log("Matched by: ", target2)
	t.Log("Matched res: ", match(target2))
	target3 := "前  other_contents" + "错误码:0" + " 后 another_contents  "
	t.Log("Matched by: ", target3)
	t.Log("Matched res: ", match(target3))
}

func TestCompileNOT(t *testing.T) {
	expRule := "P2&&!自建"
	t.Log("Expression rule: ", expRule)

	match, err := Compile(expRule)
	if err != nil {
		t.Error(err)
		return
	}
	target1 := "前  other_contents" + "自建" + " 后 another_contents  "
	t.Log("Matched by: ", target1)
	t.Log("Matched res: ", match(target1))
	target2 := "前  other_contents" + "P2空空空空空空" + " 后 another_contents  "
	t.Log("Matched by: ", target2)
	t.Log("Matched res: ", match(target2))
}
