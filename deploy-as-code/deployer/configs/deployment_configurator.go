package configs
import(
	"io/ioutil"
	"log"
	"fmt"

	"gopkg.in/yaml.v3"
)
type Output struct {
	Outputs          struct {
		ClusterEndpoint struct {
			Value string `json:"value"`
		} `json:"cluster_endpoint"`
		DbInstanceEndpoint struct {
			Value string `json:"value"`
		} `json:"db_instance_endpoint"`
		DbInstanceName struct {
			Value string `json:"value"`
		} `json:"db_instance_name"`
		DbInstancePort struct {
			Value int    `json:"value"`
		} `json:"db_instance_port"`
		DbInstanceUsername struct {
			Value     string `json:"value"`
		} `json:"db_instance_username"`
		EsDataVolumeIds struct {
			Value []string      `json:"value"`
		} `json:"es_data_volume_ids"`
		EsMasterVolumeIds struct {
			Value []string      `json:"value"`
		} `json:"es_master_volume_ids"`
		KafkaVolIds struct {
			Value []string      `json:"value"`
		} `json:"kafka_vol_ids"`
		KubectlConfig struct {
			Value string `json:"value"`
		} `json:"kubectl_config"`
		PrivateSubnets struct {
			Value []string      `json:"value"`
		} `json:"private_subnets"`
		PublicSubnets struct {
			Value []string      `json:"value"`
		} `json:"public_subnets"`
		VpcID struct {
			Value string `json:"value"`
		} `json:"vpc_id"`
		ZookeeperVolumeIds struct {
			Value []string      `json:"value"`
		} `json:"zookeeper_volume_ids"`
	} `json:"outputs"`
}
func DeployConfig(Config map[string]interface{},kvids []string,zvids []string,esdids []string,esmids []string,modules []string){
	
	file, err := ioutil.ReadFile("DIGIT-DevOps/config-as-code/environments/egov-demo.yaml")
    if err != nil {
        log.Printf("%v",err)
    }
	var data map[string]interface{}
	err = yaml.Unmarshal(file,&data)
	if err!=nil{
		log.Printf("%v",err)
	}
	for i:= range data{
		if i == "global"{
			Global:= data[i].(map[string]interface{})
			for j := range Global{
				if j== "domain"{
					Global[j]= Config["Domain"]
					fmt.Println(Global[j])
				}
			}
		}
		if i =="cluster-configs"{
			// fmt.Println("found cluster-configs")
			ClusterConfigs:=data[i].(map[string]interface{})
			for j:=range ClusterConfigs{
				if j == "configmaps"{
					// fmt.Println("found configmaps")
					Configmaps:= ClusterConfigs[j].(map[string]interface{})
					for k := range Configmaps{
						if k== "egov-config"{
							// fmt.Println("found egov-config")
							EgovConfig:= Configmaps[k].(map[string]interface{})
							for l:= range EgovConfig{
								if l=="data"{
									// fmt.Println("found data")
									Data:= EgovConfig[l].(map[string]interface{})
									for m := range Data{
										if m=="db-host"{
											Data[m]=""
											// fmt.Println(Data[m])
										}
										if m=="db-name"{
											
										}
										if m=="db-url"{
											
										} 
										if m=="domain"{
											
										}
										if m=="egov-services-fqdn-name"{
											
										}
										if m=="s3-assets-bucket"{
											
										}
										if m=="es-host"{
											
										}
										if m=="es-indexer-host"{
											
										}
										if m=="flyway-locations"{
											
										} 
										if m=="kafka-brokers"{
											
										}
										if m=="kafka-infra-brokers"{
											
										}
										if m=="logging-level-jdbc"{
											
										}
										if m=="mobile-validation-workaround"{
											
										}
										if m=="serializers-timezone-in-ist"{
											
										}
										if m=="server-tomcat-max-connections"{
											
										} 
										if m=="server-tomcat-max-threads"{
											
										}
										if m=="sms-enabled"{
											
										}
										if m=="spring-datasource-tomcat-initialSize"{
											
										}
										if m=="spring-datasource-tomcat-max-active"{
											
										} 
										if m=="spring-jpa-show-sql"{
											
										}
										if m=="timezone"{
											
										}
										if m=="tracer-errors-provideexceptionindetails"{
											
										}
									}
								}
							}
						}

					}
				}
			}
		}
		if i == "egov-mdms-service"||i=="egov-indexer"||i=="egov-persister"||i=="egov-data-uploader"||i=="egov-searcher"||i=="dashboard-analytics"||i=="dashboard-ingest"||i=="report"||i=="pdf-service"{
			// fmt.Println("in mdms")
			Service := data[i].(map[string]interface{})
			for j:=range Service{
				if j=="search-yaml-path"{

				}
				if j=="config-schema-paths"{

				}
				if j=="replicas"{

				}
				if j=="mdms-path"{

				}
				if j=="heap"{

				}
				if j=="memory_limits"{

				}
				if j=="mdms-path"{

				}
				if j=="persist-yml-path"{

				}
				if j== "initContainers"{
					// fmt.Println("in init")
					InitContainers:= Service[j].(map[string]interface{})
					for k:= range InitContainers{
						if k=="gitSync"{
							// fmt.Println("in git sync")
							GitSync:= InitContainers[k].(map[string]interface{})
							for l:=range GitSync{
								if l=="branch"{
									GitSync[l]=""
								}
								if l=="repo"{

								}
							}
						}
					}
				}
				if j=="mdms-folder"{

				}
				if j=="masters-config-url"{

				}
				if j=="java-args"{

				}
				if j=="egov-indexer-yaml-repo-path"{

				}
			}
		}
		if i=="cert-manager"{
			CertManager:=data[i].(map[string]interface{})
			for j:= range CertManager{
				if j=="email"{
					CertManager[j]=""
				}
			}
		}
		if i=="kafka-v2"{
			KafkaV2:=data[i].(map[string]interface{})
			for j:= range Kafkav2{
				if j=="persistence"{
					Persistence:=Kafkav2[j].(map[string]interface{})
					for k:= range Persistence{
						if k=="aws"{
							Aws:=Persistence[k].([]interface{})
							N:=0
							for l:= range Aws{
								Volume:=Aws[l].(map[string]interface{})
								for m:= range Volume{
									if m=="volumeId"&&N==l{
										Volume[m]=kvids[l]
									}
									if m=="zone"{

									}
								}
								N++;
								
							}
						}
					}
				}
			}
		}
		if i=="zookeeper-v2"{
			ZookeeperV2:=data[i].(map[string]interface{})
			for j:= range ZookeeperV2{
				if j=="persistence"{
					Persistence:=ZookeeperV2[j].(map[string]interface{})
					for k:= range Persistence{
						if k=="aws"{
							Aws:=Persistence[k].([]interface{})
							N:=0
							for l:= range Aws{
								Volume:=Aws[l].(map[string]interface{})
								for m:= range Volume{
									if m=="volumeId"&&N==l{
										Volume[m]=zvids[l]
									}
									if m=="zone"{
										
									}
								}
								N++;
								
							}
						}
					}
				}
			}
		}
		if i=="elasticsearch-data-v1"{
			ElasticsearchDataV1:=data[i].(map[string]interface{})
			for j:= range ElasticsearchDataV1{
				if j=="persistence"{
					Persistence:=ElasticsearchDataV1[j].(map[string]interface{})
					for k:= range Persistence{
						if k=="aws"{
							Aws:=Persistence[k].([]interface{})
							N:=0
							for l:= range Aws{
								NesteM:=Aws[l].(map[string]interface{})
								for m:= range NesteM{
									if m=="volumeId"&&N==l{
										NesteM[m]=esdids[l]
									}
									if m=="zone"{
										
									}
								}
								N++;
								
							}
						}
					}
				}
			}
		}
		if i=="elasticsearch-master-v1"{
			NestedMap:=data[i].(map[string]interface{})
			for j:= range NestedMap{
				if j=="persistence"{
					nest:=NestedMap[j].(map[string]interface{})
					for k:= range nest{
						if k=="aws"{
							Neste:=nest[k].([]interface{})
							N:=0
							for l:= range Neste{
								NesteM:=Neste[l].(map[string]interface{})
								for m:= range NesteM{
									if m=="volumeId"&&N==l{
										NesteM[m]=esmids[l]
									}
									if m=="zone"{
										
									}
								}
								N++;
								
							}
						}
					}
				}
			}
		}
		if i=="employee"{
			// delete(data,"employee")
			NestedMap:=data[i].(map[string]interface{})
			for j := range NestedMap{
				if j=="dashboard-url"{

				}
				if j=="custom-js-injection"{

				}
			}
		}
		if i=="citizen"{
			NestedMap:=data[i].(map[string]interface{})
			for j := range NestedMap{
				if j=="custom-js-injection"{

				}
			}
		}
		if i=="digit-ui"{
			delete(data,"employee")
			NestedMap:=data[i].(map[string]interface{})
			for j := range NestedMap{
				if j=="custom-js-injection"{}
			}
		}
		if i=="egov-filestore"{
			NestedMap:=data[i].(map[string]interface{})
			for j:=range NestedMap{
				if j=="volume"{

				}
				if j=="is-bucket-fixed"{

				}
				if j=="minio.url"{

				}
				if j=="aws.s3.url"{

				}
				if j=="is-s3-enabled"{

				}
				if j=="minio-enabled"{

				}
				if j=="allowed-file-formats-map"{

				}
				if j=="llowed-file-formats"{

				}
				if j=="filestore-url-validity"{

				}
				if j=="fixed-bucketname"{

				}
			}
		}
		if i=="egov-notification-sms"{
			NestedMap:=data[i].(map[string]interface{})
			for j:= range NestedMap{
				if j=="sms-provider-url"{

				}
				if j=="sms.provider.class"{
				
				}
				if j=="sms.provider.contentType"{
				
				}
				if j=="sms-config-map"{
				
				}
				if j=="sms-gateway-to-use"{
				
				}
				if j=="sms-sender"{
				
				}
				if j=="sms-sender-requesttype"{
				
				}
				if j=="sms-custom-config"{
				
				}
				if j=="sms-extra-req-params"{
				
				}
				if j=="sms-sender-req-param-name"{
				
				}
				if j=="sms-sender-username-req-param-name"{
				
				}
				if j=="sms-sender-password-req-param-name"{
				
				}
				if j=="sms-destination-mobile-req-param-name"{
				
				}
				if j=="sms-message-req-param-name"{
				
				}
				if j=="sms-error-codes"{
				
				}
			}
		}
		if i=="egov-user"{
			NestedMap:=data[i].(map[string]interface{})
			for j:= range NestedMap{
				if j=="heap"{

				}
				if j=="memory_limits"{
				
				}
				if j=="otp-validation"{
				
				}
				if j=="citizen-otp-enabled"{
				
				}
				if j=="employee-otp-enabled"{
				
				}
				if j=="access-token-validity"{
				
				}
				if j=="refresh-token-validity"{
				
				}
				if j=="default-password-expiry"{
				
				}
				if j=="mobile-number-validation"{
				
				}
				if j=="roles-state-level"{
				
				}
				if j=="zen-registration-withlogin"{
				
				}
				if j=="citizen-otp-fixed"{
				
				}
				if j=="citizen-otp-fixed-enabled"{
				
				}
				if j=="egov-state-level-tenant-id"{
				
				}
				if j=="decryption-abac-enabled"{
				
				}
			}
		}
		if i=="chatbot"{
			NestedMap:=data[i].(map[string]interface{})
			for j:= range NestedMap{
				if j=="kafka-topics-partition-count"{
  
				}
				if j=="kafka-topics-replication-factor"{
				  
				}
				if j=="kafka-consumer-poll-ms"{
				  
				}
				if j=="kafka-producer-linger-ms"{
				  
				}
				if j=="contact-card-whatsapp-number"{
				  
				}
				if j=="contact-card-whatsapp-name"{
				  
				}
				if j=="valuefirst-whatsapp-number"{
				  
				}
				if j=="valuefirst-notification-assigned-templateid"{
				  
				}
				if j=="valuefirst-notification-resolved-templateid"{
				  
				}
				if j=="valuefirst-notification-rejected-templateid"{
				  
				}
				if j=="valuefirst-notification-reassigned-templateid"{
				  
				}
				if j=="valuefirst-notification-commented-templateid"{
				  
				}
				if j=="valuefirst-notification-welcome-templateid"{
				  
				}
				if j=="valuefirst-notification-root-templateid"{
				  
				}
				if j=="valuefirst-send-message-url"{
				  
				}
				if j=="user-service-chatbot-citizen-passwrord"{
				  
				}
			}
		}
		if i=="bpa-services"{
			NestedMap:=data[i].(map[string]interface{})
			for j:= range NestedMap{
				if j=="memory_limits"{

				}
				if j=="java-args"{
				
				}
				if j=="java-debug"{
				
				}
				if j=="tracing-enabled"{
				
				}
				if j=="egov.idgen.bpa.applicationNum.format"{
				
				}
			}
		}
		if i=="bpa-calculator"{
			NestedMap:=data[i].(map[string]interface{})
			for j:= range NestedMap{
				if j=="memory_limits"{

				}
				if j=="java-args"{
				
				}
				if j=="java-debug"{
				
				}
				if j=="tracing-enabled"{
				
				}
			}
		}
		if i=="ws-services"{
			NestedMap:=data[i].(map[string]interface{})
			for j:= range NestedMap{
				if j=="wcid-format"{

				}
			}
		}
		if i=="sw-services"{
			NestedMap:=data[i].(map[string]interface{})
			for j:= range NestedMap{
				if j=="scid-format"{

				}
			}
		}
		if i=="egov-pg-service"{
			NestedMap:=data[i].(map[string]interface{})
			for j:= range NestedMap{
				if j=="axis"{

				}
			}
		}
		if i=="report"{
			NestedMap:=data[i].(map[string]interface{})
			for j:=range NestedMap{
				if j=="heap"{
					
				}
				if j=="tracing-enabled"{
					
				}
				if j=="spring-datasource-tomcat-max-active"{
					
				}
				if j=="initContainers"{
					NestedM:= NestedMap[j].(map[string]interface{})
					for k := range NestedM{
						if k=="gitSync"{
							Neste:=NestedM[k].(map[string]interface{})
							for l:= range Neste{
								if l=="repo"{
							
								}
								if l=="branch"{
									
								}
							}
						}
					}
				}    
				if j=="report-locationsfile-path"{
					
				}
			}
		}
		if i=="pdf-service"{
			NestedMap:=data[i].(map[string]interface{})
			for j:=range NestedMap{
				if j=="initContainers"{
					NestedM:= NestedMap[j].(map[string]interface{})
					for k := range NestedM{
						if k=="gitSync"{
							Neste:=NestedM[k].(map[string]interface{})
							for l:= range Neste{
								if l=="repo"{
							
								}
								if l=="branch"{
									
								}
							}
						}
					}
				}
				if j=="data-config-urls"{

				}
				if j=="format-config-urls"{

				}
				   
			}
		}
		if i=="egf-master"{
			NestedMap:=data[i].(map[string]interface{})
			for j:= range NestedMap{
				if j=="db-url"{

				}
				if j=="memory_limits"{
				
				}
				if j=="heap"{
				
				}
				
			}
		}
		if i=="egov-custom-consumer"{
			NestedMap:=data[i].(map[string]interface{})
			for j:= range NestedMap{
				if j=="erp-host"{

				}
			}
		}
		if i=="egov-apportion-service"{
			NestedMap:=data[i].(map[string]interface{})
			for j:= range NestedMap{
				if j=="memory_limits"{

				}
				if j=="heap"{

				}
			}
		}
		if i=="redoc"{
			NestedMap:=data[i].(map[string]interface{})
			for j:= range NestedMap{
				if j=="replicas"{

				}
				if j=="images"{

				}
				if j=="service_type"{

				}
			}
		}
		if i=="redoc"{
			NestedMap:=data[i].(map[string]interface{})
			for j:= range NestedMap{
				if j=="images"{

				}
				if j=="replicas"{
				
				}
				if j=="default-backend-service"{
				
				}
				if j=="namespace"{
				
				}
				if j=="cert-issuer"{
				
				}
				if j=="ssl-protocols"{
				
				}
				if j=="ssl-ciphers"{
				
				}
				if j=="ssl-ecdh-curve"{
				
				}
			}
		}
		if i=="cert-manager"{
			NestedMap:=data[i].(map[string]interface{})
			for j:= range NestedMap{
				if j=="email"{

				}
			}
		}
		if i=="zuul"{
			NestedMap:=data[i].(map[string]interface{})
			for j:= range NestedMap{
				if j=="replicas"{

				}
				if j=="custom-filter-property"{
				
				}
				if j=="tracing-enabled"{
				
				}
				if j=="heap"{
				
				}
				if j=="server-tomcat-max-threads"{
				
				}
				if j=="server-tomcat-max-connections"{
				
				}
				if j=="egov-open-endpoints-whitelist"{
				
				}
				if j=="egov-mixed-mode-endpoints-whitelist"{
				
				}
			}
		}
		if i=="collection-services"{
			NestedMap:=data[i].(map[string]interface{})
			for j:= range NestedMap{
				if j=="receiptnumber-servicebased"{

				}
				if j=="receipt-search-paginate"{
				
				}
				if j=="receipt-search-defaultsize"{
				
				}
				if j=="user-create-enabled"{
				
				}
			}
		}
		if i=="collection-receipt-voucher-consumer"{
			NestedMap:=data[i].(map[string]interface{})
			for j:= range NestedMap{
				if j=="jalandhar-erp-host"{

				}
				if j=="mohali-erp-host"{
				
				}
				if j=="nayagaon-erp-host"{
				
				}
				if j=="amritsar-erp-host"{
				
				}
				if j=="kharar-erp-host"{
				
				}
				if j=="zirakpur-erp-host"{
				
				}
			}
		}
		if i=="finance-collections-voucher-consumer"{
			NestedMap:=data[i].(map[string]interface{})
			for j:= range NestedMap{
				if j=="erp-env-name"{

				}
				if j=="erp-domain-name"{

				}
			}
		}
		if i=="rainmaker-pgr"{
			NestedMap:=data[i].(map[string]interface{})
			for j:= range NestedMap{
				if j=="notification-sms-enabled"{

				}
				if j=="notification-email-enabled"{
				
				}
				if j=="new-complaint-enabled"{
				
				}
				if j=="reassign-complaint-enabled"{
				
				}
				if j=="reopen-complaint-enabled"{
				
				}
				if j=="comment-by-employee-notif-enabled"{
				
				}
				if j=="notification-allowed-status"{
				
				}
			}
		}
		if i=="pt-services-v2"{
			NestedMap:=data[i].(map[string]interface{})
			for j:= range NestedMap{
				if j=="pt-userevents-pay-link"{

				}
			}
		}
		if i=="pt-calculator-v2"{
			NestedMap:=data[i].(map[string]interface{})
			for j:= range NestedMap{
				if j=="logging-level"{

				}
			}
		}
		if i=="tl-services"{
			NestedMap:=data[i].(map[string]interface{})
			for j:= range NestedMap{
				if j=="heap"{

				}
				if j=="memory_limits"{
				
				}
				if j=="java-args"{
				
				}
				if j=="tl-application-num-format"{
				
				}
				if j=="tl-license-num-format"{
				
				}
				if j=="tl-userevents-pay-link"{
				
				}
				if j=="tl-payment-topic-name"{
				
				}
				if j=="host-link"{
				
				}
				if j=="pdf-link"{
				
				}
				if j=="tl-search-default-limit"{
				
				}
			}
		}
		if i=="egov-hrms"{
			NestedMap:=data[i].(map[string]interface{})
			for j:= range NestedMap{
				if j=="java-args"{

				}
				if j=="heap"{
				
				}
				if j=="employee-applink"{
				
				}
			}
		}
		if i=="egov-weekly-impact-notifier"{
			NestedMap:=data[i].(map[string]interface{})
			for j:= range NestedMap{
				if j=="mail-to-address"{

				}
				if j=="mail-interval-in-secs"{
				
				}
				if j=="schedule"{
				
				}
			}
		}
		if i=="kafka-config"{
			NestedMap:=data[i].(map[string]interface{})
			for j:= range NestedMap{
				if j=="topics"{

				}
				if j=="zookeeper-connect"{
				
				}
				if j=="kafka-brokers"{
				
				}
			}
		}
		if i=="logging-config"{
			NestedMap:=data[i].(map[string]interface{})
			for j:= range NestedMap{
				if j=="es-host"{

				}
				if j=="es-port"{
				
				}
			}
		}
		if i=="jaeger-config"{
			NestedMap:=data[i].(map[string]interface{})
			for j:= range NestedMap{
				if j=="host"{

				}
				if j=="port"{
				
				}
				if j=="sampler-type"{
				
				}
				if j=="sampler-param"{
				
				}
				if j=="sampling-strategies"{
				
				}
			}
		}
		if i=="redis"{
			NestedMap:=data[i].(map[string]interface{})
			for j:= range NestedMap{
				if j=="replicas"{

				}
				if j=="images"{
				
				}
			}
		}
		if i=="playground"{
			NestedMap:=data[i].(map[string]interface{})
			for j:= range NestedMap{
				if j=="replicas"{

				}
				if j=="images"{
				
				}
			}
		}
		if i=="fluent-bit"{
			NestedMap:=data[i].(map[string]interface{})
			for j:= range NestedMap{
				if j=="images"{

				}
				if j=="egov-services-log-topic"{
				
				}
				if j=="egov-infra-log-topic"{

				}
			}
		}
		if i=="egov-workflow-v2"{
			NestedMap:=data[i].(map[string]interface{})
			for j:= range NestedMap{
				if j=="logging-level"{

				}
				if j=="java-args"{
				
				}
				if j=="heap"{
				
				}
				if j=="workflow-statelevel"{
				
				}
				if j=="host-link"{
				
				}
				if j=="pdf-link"{
				
				}
			}
		}
	}
	n:=len(modules)
	for _,i:= range modules{
		if i=="m_pgr"&&n==1{
			
		}
		if i=="m_property-tax"&&n==1{
			
		}
		if i=="m_sewerage"&&n==1{
			
		}
		if i=="m_bpa"&&n==1{
			
		}
		if i=="m_trade-license"&&n==1{
			
		}
		if i=="m_firenoc"&&n==1{
			
		}
		if i=="m_water-service"&&n==1{
			
		}
		if i=="m_dss"&&n==1{
			
		}
		if i=="m_fsm"&&n==1{
			
		}
		if i=="m_echallan"&&n==1{
			
		}
		if i=="m_edcr"&&n==1{
			
		}
		if i=="m_finance"&&n==1{
			
		}
		if i=="m_pgr"&&i=="m_property-tax"&&n==2{
			
		}
		if i=="m_pgr"&&i=="m_sewerage"&&n==2{
	
		}
		if i=="m_pgr"&&i=="m_bpa"&&n==2{
			
		}
		if i=="m_pgr"&&i=="m_trade-license"&&n==2{
			
		}
		if i=="m_pgr"&&i=="m_firenoc"&&n==2{
			
		}
		if i=="m_pgr"&&i=="m_water-service"&&n==2{
			
		}
		if i=="m_pgr"&&i=="m_dss"&&n==2{
			
		}
		if i=="m_pgr"&&i=="m_fsm"&&n==2{
			
		}
		if i=="m_pgr"&&i=="m_echallan"&&n==2{
			
		}
		if i=="m_pgr"&&i=="m_edcr"&&n==2{
			
		}
		if i=="m_pgr"&&i=="m_finance"&&n==2{
			
		}
		if i=="m_property-tax"&&i=="m_sewerage"&&n==2{
	
		}
		if i=="m_property-tax"&&i=="m_bpa"&&n==2{
			
		}
		if i=="m_property-tax"&&i=="m_trade-license"&&n==2{
			
		}
		if i=="m_property-tax"&&i=="m_firenoc"&&n==2{
			
		}
		if i=="m_property-tax"&&i=="m_water-service"&&n==2{
			
		}
		if i=="m_property-tax"&&i=="m_dss"&&n==2{
			
		}
		if i=="m_property-tax"&&i=="m_fsm"&&n==2{
			
		}
		if i=="m_property-tax"&&i=="m_echallan"&&n==2{
			
		}
		if i=="m_property-tax"&&i=="m_edcr"&&n==2{
			
		}
		if i=="m_property-tax"&&i=="m_finance"&&n==2{
			
		}
		if i=="m_sewerage"&&i=="m_bpa"&&n==2{
	
		}
		if i=="m_sewerage"&&i=="m_trade-license"&&n==2{
			
		}
		if i=="m_sewerage"&&i=="m_firenoc"&&n==2{
			
		}
		if i=="m_sewerage"&&i=="m_water-service"&&n==2{
			
		}
		if i=="m_sewerage"&&i=="m_dss"&&n==2{
			
		}
		if i=="m_sewerage"&&i=="m_fsm"&&n==2{
			
		}
		if i=="m_sewerage"&&i=="m_echallan"&&n==2{
			
		}
		if i=="m_sewerage"&&i=="m_edcr"&&n==2{
			
		}
		if i=="m_sewerage"&&i=="m_finance"&&n==2{
			
		}
		if i=="m_bpa"&&i=="m_trade-license"&&n==2{
	
		}
		if i=="m_bpa"&&i=="m_firenoc"&&n==2{
			
		}
		if i=="m_bpa"&&i=="m_water-service"&&n==2{
			
		}
		if i=="m_bpa"&&i=="m_dss"&&n==2{
			
		}
		if i=="m_bpa"&&i=="m_fsm"&&n==2{
			
		}
		if i=="m_bpa"&&i=="m_echallan"&&n==2{
			
		}
		if i=="m_bpa"&&i=="m_edcr"&&n==2{
			
		}
		if i=="m_bpa"&&i=="m_finance"&&n==2{
			
		}
		if i=="m_trade-license"&&i=="m_firenoc"&&n==2{
	
		}
		if i=="m_trade-license"&&i=="m_water-service"&&n==2{
			
		}
		if i=="m_trade-license"&&i=="m_dss"&&n==2{
			
		}
		if i=="m_trade-license"&&i=="m_fsm"&&n==2{
			
		}
		if i=="m_trade-license"&&i=="m_echallan"&&n==2{
			
		}
		if i=="m_trade-license"&&i=="m_edcr"&&n==2{
			
		}
		if i=="m_trade-license"&&i=="m_finance"&&n==2{
			
		}
		if i=="m_firenoc"&&i=="m_water-service"&&n==2{
	
		}
		if i=="m_firenoc"&&i=="m_dss"&&n==2{
			
		}
		if i=="m_firenoc"&&i=="m_fsm"&&n==2{
			
		}
		if i=="m_firenoc"&&i=="m_echallan"&&n==2{
			
		}
		if i=="m_firenoc"&&i=="m_edcr"&&n==2{
			
		}
		if i=="m_firenoc"&&i=="m_finance"&&n==2{
			
		}
		if i=="m_water-service"&&i=="m_dss"&&n==2{
	
		}
		if i=="m_water-service"&&i=="m_fsm"&&n==2{
			
		}
		if i=="m_water-service"&&i=="m_echallan"&&n==2{
			
		}
		if i=="m_water-service"&&i=="m_edcr"&&n==2{
			
		}
		if i=="m_water-service"&&i=="m_finance"&&n==2{
			
		}
		if i=="m_dss"&&i=="m_fsm"&&n==2{
	
		}
		if i=="m_dss"&&i=="m_echallan"&&n==2{
			
		}
		if i=="m_dss"&&i=="m_edcr"&&n==2{
			
		}
		if i=="m_dss"&&i=="m_finance"&&n==2{
			
		}
		if i=="m_fsm"&&i=="m_echallan"&&n==2{
	
		}
		if i=="m_fsm"&&i=="m_edcr"&&n==2{
			
		}
		if i=="m_fsm"&&i=="m_finance"&&n==2{
			
		}
		if i=="m_echallan"&&i=="m_edcr"&&n==2{
	
		}
		if i=="m_echallan"&&i=="m_finance"&&n==2{
			
		}
		if i=="m_edcr"&&i=="m_finance"&&n==2{
	
		}		
		if i=="m_pgr"&&i=="m_property-tax"&&i=="m_sewerage"&&n==3{
	
		}
		if i=="m_bpa"&&i=="m_property-tax"&&i=="m_sewerage"&&n==3{
			
		}
		if i=="m_trade-license"&&i=="m_property-tax"&&i=="m_sewerage"&&n==3{
			
		}
		if i=="m_firenoc"&&i=="m_property-tax"&&i=="m_sewerage"&&n==3{
			
		}
		if i=="m_water-service"&&i=="m_property-tax"&&i=="m_sewerage"&&n==3{
			
		}
		if i=="m_dss"&&i=="m_property-tax"&&i=="m_sewerage"&&n==3{
			
		}
		if i=="m_fsm"&&i=="m_property-tax"&&i=="m_sewerage"&&n==3{
			
		}
		if i=="m_echallan"&&i=="m_property-tax"&&i=="m_sewerage"&&n==3{
			
		}
		if i=="m_edcr"&&i=="m_property-tax"&&i=="m_sewerage"&&n==3{
			
		}
		if i=="m_finance"&&i=="m_property-tax"&&i=="m_sewerage"&&n==3{
			
		}
		if i=="m_pgr"&&i=="m_bpa"&&i=="m_sewerage"&&n==3{
			
		}
		if i=="m_pgr"&&i=="m_trade-license"&&i=="m_sewerage"&&n==3{
			
		}
		if i=="m_pgr"&&i=="m_firenoc"&&i=="m_sewerage"&&n==3{
			
		}
		if i=="m_pgr"&&i=="m_water-service"&&i=="m_sewerage"&&n==3{
			
		}
		if i=="m_pgr"&&i=="m_dss"&&i=="m_sewerage"&&n==3{
			
		}
		if i=="m_pgr"&&i=="m_fsm"&&i=="m_sewerage"&&n==3{
			
		}
		if i=="m_pgr"&&i=="m_echallan"&&i=="m_sewerage"&&n==3{
			
		}
		if i=="m_pgr"&&i=="m_edcr"&&i=="m_sewerage"&&n==3{
			
		}
		if i=="m_pgr"&&i=="m_finance"&&i=="m_sewerage"&&n==3{
			
		}
		if i=="m_pgr"&&i=="m_property-tax"&&i=="m_bpa"&&n==3{
			
		}
		if i=="m_pgr"&&i=="m_property-tax"&&i=="m_trade-license"&&n==3{
			
		}
		if i=="m_pgr"&&i=="m_property-tax"&&i=="m_firenoc"&&n==3{
			
		}
		if i=="m_pgr"&&i=="m_property-tax"&&i=="m_water-service"&&n==3{
			
		}
		if i=="m_pgr"&&i=="m_property-tax"&&i=="m_dss"&&n==3{
			
		}
		if i=="m_pgr"&&i=="m_property-tax"&&i=="m_fsm"&&n==3{
			
		}
		if i=="m_pgr"&&i=="m_property-tax"&&i=="m_echallan"&&n==3{
			
		}
		if i=="m_pgr"&&i=="m_property-tax"&&i=="m_edcr"&&n==3{
			
		}
		if i=="m_pgr"&&i=="m_property-tax"&&i=="m_finance"&&n==3{
			
		}
		if i=="m_pgr"&&i=="m_property-tax"&&i=="m_finance"&&n==3{
			
		}
		if i=="m_bpa"&&i=="m_trade-license"&&i=="m_firenoc"&&n==3{
			
		}
		if i=="m_pgr"&&i=="m_trade-license"&&i=="m_firenoc"&&n==3{
			
		}
		if i=="m_property-tax"&&i=="m_trade-license"&&i=="m_firenoc"&&n==3{
			
		}
		if i=="m_sewerage"&&i=="m_trade-license"&&i=="m_firenoc"&&n==3{
			
		}
		if i=="m_water-service"&&i=="m_trade-license"&&i=="m_firenoc"&&n==3{
			
		}
		if i=="m_dss"&&i=="m_trade-license"&&i=="m_firenoc"&&n==3{
			
		}
		if i=="m_fsm"&&i=="m_trade-license"&&i=="m_firenoc"&&n==3{
			
		}
		if i=="m_echallan"&&i=="m_trade-license"&&i=="m_firenoc"&&n==3{
			
		}
		if i=="m_edcr"&&i=="m_trade-license"&&i=="m_firenoc"&&n==3{
			
		}
		if i=="m_finance"&&i=="m_trade-license"&&i=="m_firenoc"&&n==3{
			
		}
		if i=="m_bpa"&&i=="m_pgr"&&i=="m_firenoc"&&n==3{
			
		}
		if i=="m_bpa"&&i=="m_property-tax"&&i=="m_firenoc"&&n==3{
			
		}
		if i=="m_bpa"&&i=="m_sewerage"&&i=="m_firenoc"&&n==3{
			
		}
		if i=="m_bpa"&&i=="m_water-service"&&i=="m_firenoc"&&n==3{
			
		}
		if i=="m_bpa"&&i=="m_dss"&&i=="m_firenoc"&&n==3{
			
		}
		if i=="m_bpa"&&i=="m_fsm"&&i=="m_firenoc"&&n==3{
			
		}
		if i=="m_bpa"&&i=="m_echallan"&&i=="m_firenoc"&&n==3{
			
		}
		if i=="m_bpa"&&i=="m_edcr"&&i=="m_firenoc"&&n==3{
			
		}
		if i=="m_bpa"&&i=="m_finance"&&i=="m_firenoc"&&n==3{
			
		}
		if i=="m_bpa"&&i=="m_trade-license"&&i=="m_pgr"&&n==3{
			
		}
		if i=="m_bpa"&&i=="m_trade-license"&&i=="m_property-tax"&&n==3{
			
		}
		if i=="m_bpa"&&i=="m_trade-license"&&i=="m_sewerage"&&n==3{
			
		}
		if i=="m_bpa"&&i=="m_trade-license"&&i=="m_water-service"&&n==3{
			
		}
		if i=="m_bpa"&&i=="m_trade-license"&&i=="m_dss"&&n==3{
			
		}
		if i=="m_bpa"&&i=="m_trade-license"&&i=="m_fsm"&&n==3{
			
		}
		if i=="m_bpa"&&i=="m_trade-license"&&i=="m_echallan"&&n==3{
			
		}
		if i=="m_bpa"&&i=="m_trade-license"&&i=="m_edcr"&&n==3{
			
		}
		if i=="m_bpa"&&i=="m_trade-license"&&i=="m_finance"&&n==3{
			
		}
		if i=="m_water-service"&&i=="m_dss"&&i=="m_fsm"&&n==3{
			
		}
		if i=="m_pgr"&&i=="m_dss"&&i=="m_fsm"&&n==3{
			
		}
		if i=="m_property-tax"&&i=="m_dss"&&i=="m_fsm"&&n==3{
			
		}
		if i=="m_sewerage"&&i=="m_dss"&&i=="m_fsm"&&n==3{
			
		}
		if i=="m_bpa"&&i=="m_dss"&&i=="m_fsm"&&n==3{
			
		}
		if i=="m_trade-license"&&i=="m_dss"&&i=="m_fsm"&&n==3{
			
		}
		if i=="m_firenoc"&&i=="m_dss"&&i=="m_fsm"&&n==3{
			
		}
		if i=="m_echallan"&&i=="m_dss"&&i=="m_fsm"&&n==3{
			
		}
		if i=="m_edcr"&&i=="m_dss"&&i=="m_fsm"&&n==3{
			
		}
		if i=="m_finance"&&i=="m_dss"&&i=="m_fsm"&&n==3{
			
		}
		if i=="m_water-service"&&i=="m_pgr"&&i=="m_fsm"&&n==3{
			
		}
		if i=="m_water-service"&&i=="m_property-tax"&&i=="m_fsm"&&n==3{
			
		}
		if i=="m_water-service"&&i=="m_sewerage"&&i=="m_fsm"&&n==3{
			
		}
		if i=="m_water-service"&&i=="m_bpa"&&i=="m_fsm"&&n==3{
			
		}
		if i=="m_water-service"&&i=="m_trade-license"&&i=="m_fsm"&&n==3{
			
		}
		if i=="m_water-service"&&i=="m_firenoc"&&i=="m_fsm"&&n==3{
			
		}
		if i=="m_water-service"&&i=="m_echallan"&&i=="m_fsm"&&n==3{
			
		}
		if i=="m_water-service"&&i=="m_edcr"&&i=="m_fsm"&&n==3{
			
		}
		if i=="m_water-service"&&i=="m_finance"&&i=="m_fsm"&&n==3{
			
		}
		if i=="m_water-service"&&i=="m_dss"&&i=="m_pgr"&&n==3{
			
		}
		if i=="m_water-service"&&i=="m_dss"&&i=="m_property-tax"&&n==3{
			
		}
		if i=="m_water-service"&&i=="m_dss"&&i=="m_sewerage"&&n==3{
			
		}
		if i=="m_water-service"&&i=="m_dss"&&i=="m_bpa"&&n==3{
			
		}
		if i=="m_water-service"&&i=="m_dss"&&i=="m_trade-license"&&n==3{
			
		}
		if i=="m_water-service"&&i=="m_dss"&&i=="m_firenoc"&&n==3{
			
		}
		if i=="m_water-service"&&i=="m_dss"&&i=="m_fsm"&&n==3{
			
		}
		if i=="m_water-service"&&i=="m_dss"&&i=="m_echallan"&&n==3{
			
		}
		if i=="m_water-service"&&i=="m_dss"&&i=="m_edcr"&&n==3{
			
		}
		if i=="m_water-service"&&i=="m_dss"&&i=="m_finance"&&n==3{
			
		}
		if i=="m_echallan"&&i=="m_edcr"&&i=="m_finance"&&n==3{
			
		}
		if i=="m_pgr"&&i=="m_edcr"&&i=="m_finance"&&n==3{
			
		}
		if i=="m_property-tax"&&i=="m_edcr"&&i=="m_finance"&&n==3{
			
		}
		if i=="m_sewerage"&&i=="m_edcr"&&i=="m_finance"&&n==3{
			
		}
		if i=="m_bpa"&&i=="m_edcr"&&i=="m_finance"&&n==3{
			
		}
		if i=="m_trade-license"&&i=="m_edcr"&&i=="m_finance"&&n==3{
			
		}
		if i=="m_firenoc"&&i=="m_edcr"&&i=="m_finance"&&n==3{
			
		}
		if i=="m_water-service"&&i=="m_edcr"&&i=="m_finance"&&n==3{
			
		}
		if i=="m_dss"&&i=="m_edcr"&&i=="m_finance"&&n==3{
			
		}
		if i=="m_fsm"&&i=="m_edcr"&&i=="m_finance"&&n==3{
			
		}
		if i=="m_echallan"&&i=="m_property-tax"&&i=="m_finance"&&n==3{
			
		}
		if i=="m_echallan"&&i=="m_sewerage"&&i=="m_finance"&&n==3{
			
		}
		if i=="m_echallan"&&i=="m_bpa"&&i=="m_finance"&&n==3{
			
		}
		if i=="m_echallan"&&i=="m_trade-license"&&i=="m_finance"&&n==3{
			
		}
		if i=="m_echallan"&&i=="m_firenoc"&&i=="m_finance"&&n==3{
			
		}
		if i=="m_echallan"&&i=="m_water-service"&&i=="m_finance"&&n==3{
			
		}
		if i=="m_echallan"&&i=="m_dss"&&i=="m_finance"&&n==3{
			
		}
		if i=="m_echallan"&&i=="m_fsm"&&i=="m_finance"&&n==3{
			
		}
		if i=="m_echallan"&&i=="m_edcr"&&i=="m_pgr"&&n==3{
			
		}
		if i=="m_echallan"&&i=="m_edcr"&&i=="m_property-tax"&&n==3{
			
		}
		if i=="m_echallan"&&i=="m_edcr"&&i=="m_sewerage"&&n==3{
			
		}
		if i=="m_echallan"&&i=="m_edcr"&&i=="m_bpa"&&n==3{
			
		}
		if i=="m_echallan"&&i=="m_edcr"&&i=="m_trade-license"&&n==3{
			
		}
		if i=="m_echallan"&&i=="m_edcr"&&i=="m_firenoc"&&n==3{
			
		}
		if i=="m_echallan"&&i=="m_edcr"&&i=="m_water-service"&&n==3{
			
		}
		if i=="m_echallan"&&i=="m_edcr"&&i=="m_dss"&&n==3{
			
		}
		if i=="m_echallan"&&i=="m_edcr"&&i=="m_fsm"&&n==3{
			
		}
		if i=="m_echallan"&&i=="m_edcr"&&i=="m_finance"&&n==3{
			
		}
	}
	newfile,err := yaml.Marshal(&data)
	if err!=nil{
		log.Printf("%v",err)

	}
	err=ioutil.WriteFile("new_cr.yaml",newfile,0644)
	if err!=nil{
		log.Printf("%v",err)
	}
}