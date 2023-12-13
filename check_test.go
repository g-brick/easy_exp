package easy_exp

import (
	"encoding/json"
	"testing"
)

type exampleRule struct {
	Keyword string `json:"keyword"`
}

var jsonExampleString = `[
    {
      "keyword": "golang1||golang2"
    },
    {
      "keyword": "golang2||golang11"
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
      "keyword": "golang2"
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
      "keyword": "\"golang2\"&&\"上海-自建\"||\"golang2\"&&\"广州-自建\"||\"golang2\"&&\"北京-自建\""
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
		check, err := Compile(v.Keyword)
		if err != nil {
			t.Error(err)
			continue
		}
		target := "前  other_contents" + v.Keyword + " 后 another_contents  "
		t.Log("Matched by: ", target)
		t.Log("Matched res: ", check(target))
	}
}

func TestCompileMix1(t *testing.T) {
	expRule := "\"golang2\"&&\"上海-自建\"||\"golang2\"&&\"广州-自建\"||\"golang2\"&&\"北京-自建\""
	t.Log("Expression rule: ", expRule)
	check, err := Compile(expRule)
	if err != nil {
		t.Error(err)
		return
	}
	target1 := "前  other_contents" + "广州-自建" + " 后 another_contents  " + "golang2"
	t.Log("Matched by: ", target1)
	t.Log("Matched res: ", check(target1))
	target2 := "前  other_contents" + "广州-自建" + " 后 another_contents  " + "golang1"
	t.Log("Matched by: ", target2)
	t.Log("Matched res: ", check(target2))
	target3 := "前  other_contents" + "golang2" + " 后 another_contents  "
	t.Log("Matched by: ", target3)
	t.Log("Matched res: ", check(target3))
}

func TestCompileMix2(t *testing.T) {
	expRule := "\"错误码:26904\"||\"错误码:0\"&& ( UMON||Monitor||UMonNetwork||TianLiang )"
	t.Log("Expression rule: ", expRule)

	check, err := Compile(expRule)
	if err != nil {
		t.Error(err)
		return
	}
	target1 := "前  other_contents" + "错误码:26904" + " 后 another_contents  "
	t.Log("Matched by: ", target1)
	t.Log("Matched res: ", check(target1))
	target2 := "前  other_contents" + "错误码:0" + " 后 another_contents  " + "UMonNetwork"
	t.Log("Matched by: ", target2)
	t.Log("Matched res: ", check(target2))
	target3 := "前  other_contents" + "错误码:0" + " 后 another_contents  "
	t.Log("Matched by: ", target3)
	t.Log("Matched res: ", check(target3))
}

func TestCompileNOT(t *testing.T) {
	expRule := "golang2&&!自建"
	t.Log("Expression rule: ", expRule)

	check, err := Compile(expRule)
	if err != nil {
		t.Error(err)
		return
	}
	target1 := "前  other_contents" + "自建" + " 后 another_contents  "
	t.Log("Matched by: ", target1)
	t.Log("Matched res: ", check(target1))
	target2 := "前  other_contents" + "golang2空空空空空空" + " 后 another_contents  "
	t.Log("Matched by: ", target2)
	t.Log("Matched res: ", check(target2))
	target3 := "前  other_contents" + "!自建" + " 后 another_contents  "
	t.Log("Matched by: ", target3)
	t.Log("Matched res: ", check(target3))
}

func TestCompileNormalTxt(t *testing.T) {
	expRule := "John" // If some txt includes John, it will be true
	t.Log("expRule: ", expRule)
	check, err := Compile(expRule)
	if err != nil {
		t.Error(err)
		return
	}
	target1 := "My college teacher is Mr.John"
	t.Log("Matched by: ", target1)
	t.Log("True is expected: ", check(target1))
	target2 := "My best friend is Jay"
	t.Log("Matched by: ", target2)
	t.Log("False is expected: ", check(target2))
	target3 := "Yesterday, John got sick."
	t.Log("Matched by: ", target3)
	t.Log("True is expected: ", check(target3))
}

func TestCompileNotRule(t *testing.T) {
	expRule := "!John" // If some txt includes John, it will be false
	t.Log("expRule: ", expRule)
	check, err := Compile(expRule)
	if err != nil {
		t.Error(err)
		return
	}
	target1 := "My college teacher is Mr.John"
	t.Log("Matched by: ", target1)
	t.Log("False is expected: ", check(target1))
	target2 := "My best friend is Jay"
	t.Log("Matched by: ", target2)
	t.Log("True is expected: ", check(target2))
	target3 := "Yesterday, John and Joy got sick."
	t.Log("Matched by: ", target3)
	t.Log("False is expected: ", check(target3))
}

func TestCompileAndRule(t *testing.T) {
	expRule := "\"golang1.1\"&&\"golang1.2\""
	t.Log("Expression rule: ", expRule)
	check, err := Compile(expRule)
	if err != nil {
		t.Error(err)
		return
	}
	target1 := "I has made my golang version upgrade from golang1.1 to golang1.2"
	t.Log("Matched by: ", target1)
	t.Log("True is expected: ", check(target1))
	target2 := "I has made my golang version downgrade from golang1.2 to golang1.1"
	t.Log("Matched by: ", target2)
	t.Log("True is expected: ", check(target2))
	target3 := "The version of go in this program is golang1.1" // No golang1.2 in the txt, so result is false
	t.Log("Matched by: ", target3)
	t.Log("False is expected: ", check(target3))
}

func TestCompileOrRule(t *testing.T) {
	expRule := "\"golang1.1\"||\"golang1.2\""
	t.Log("Expression rule: ", expRule)
	check, err := Compile(expRule)
	if err != nil {
		t.Error(err)
		return
	}
	target1 := "I has made my golang version upgrade from golang1.1 to golang1.2"
	t.Log("Matched by: ", target1)
	t.Log("True is expected: ", check(target1))
	target2 := "I has made my golang version downgrade from golang1.2 to golang1.1"
	t.Log("Matched by: ", target2)
	t.Log("True is expected: ", check(target2))
	target3 := "The version of go in this program is golang1.1"
	t.Log("Matched by: ", target3)
	t.Log("True is expected: ", check(target3))
	target4 := "The version of go in this program is golang1.2"
	t.Log("Matched by: ", target4)
	t.Log("True is expected: ", check(target4))
	target5 := "The version of go in this program is golang1.0" // Not matched the rule:  golang1.1 || golang1.2
	t.Log("Matched by: ", target5)
	t.Log("False is expected: ", check(target5))
}

func TestCompileMixedRule1(t *testing.T) {
	expRule := "golang1.2&&!golang1.1" // if some txt includes the golang1.2 and excludes the golang1.1, it will be true
	t.Log("Expression rule: ", expRule)

	check, err := Compile(expRule)
	if err != nil {
		t.Error(err)
		return
	}
	target1 := "I has made my golang version upgrade from golang1.1 to golang1.2"
	t.Log("Matched by: ", target1)
	t.Log("False is expected: ", check(target1))
	target2 := "I has made my golang version upgrade from golang1.2 to golang1.3"
	t.Log("Matched by: ", target2)
	t.Log("True is expected: ", check(target2))
	target3 := "The version of go in this program is golang1.1"
	t.Log("Matched by: ", target3)
	t.Log("False is expected: ", check(target3))
	target4 := "The version of go in this program is golang1.2"
	t.Log("Matched by: ", target4)
	t.Log("True is expected: ", check(target4))
	target5 := "The version of go in this program is golang1.0"
	t.Log("Matched by: ", target5)
	t.Log("False is expected: ", check(target5))
}

func TestCompileMixedRule2(t *testing.T) {
	expRule := "(\"China\"&&\"USA\")||(\"Japan\"&&\"Korea\")||(\"Asia\"&&\"North American\")"
	// expRule equals expRule2
	// expRule2 := "(China && USA) || (Japan && Korea) || (Asia && North American)"
	t.Log("Expression rule: ", expRule)
	check, err := Compile(expRule)
	if err != nil {
		t.Error(err)
		return
	}
	target1 := "I am from USA, but I love China."
	t.Log("Matched by: ", target1)
	t.Log("True is expected: ", check(target1))
	target2 := "Japan is not in Africa, and neither is South Korea."
	t.Log("Matched by: ", target2)
	t.Log("True is expected: ", check(target2))
	target3 := "The North American is smaller than Asia."
	t.Log("Matched by: ", target3)
	t.Log("True is expected: ", check(target3))
	target4 := "The USA is a part of the North American."
	t.Log("Matched by: ", target4)
	t.Log("False is expected: ", check(target4))
}

func TestCompileMixedRule3(t *testing.T) {
	expRule := "\"China\"&&\"USA\"||\"Japan\"&&\"Korea\"" // If no parenthesis in the expression, it also works.
	// expRule equals expRule2
	// expRule2 := "China && USA || Japan && Korea
	t.Log("Expression rule: ", expRule)
	check, err := Compile(expRule)
	if err != nil {
		t.Error(err)
		return
	}
	target1 := "I am from USA, but I love China."
	t.Log("Matched by: ", target1)
	t.Log("True is expected: ", check(target1))
	target2 := "I am from China, and I want to travel to Japan or Korea."
	t.Log("Matched by: ", target2)
	t.Log("True is expected: ", check(target2))
	target3 := "I am from USA, and I want to travel to Japan."
	t.Log("Matched by: ", target3)
	t.Log("False is expected: ", check(target3))
	target4 := "I am from USA, and I want to travel to Korea."
	t.Log("Matched by: ", target4)
	t.Log("False is expected: ", check(target4))
}

func TestCompileMixedRule4(t *testing.T) {
	expRule := "(!\"JAVA\"||\"Golang\")&&\"PHP\""
	// expRule equals expRule2
	// expRule2 := "(!JAVA || Golang) && PHP"
	t.Log("Expression rule: ", expRule)
	check, err := Compile(expRule)
	if err != nil {
		t.Error(err)
		return
	}
	target1 := "I am good at JAVA, not PHP."
	t.Log("Matched by: ", target1)
	t.Log("False is expected: ", check(target1))
	target2 := "I am good at Golang, not PHP."
	t.Log("Matched by: ", target2)
	t.Log("True is expected: ", check(target2))
	target3 := "I am good at PHP." // Match !JAVA && PHP
	t.Log("Matched by: ", target3)
	t.Log("True is expected: ", check(target3))
	target4 := "I am good at Golang."
	t.Log("Matched by: ", target4)
	t.Log("False is expected: ", check(target4))
}
