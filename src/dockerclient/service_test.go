package dockerclient

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"encoding/json"
	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/swarm"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateService(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"ID":"e90302"}`))
	}))
	defer server.Close()

	httpClient, err := NewHttpClient()
	assert.Nil(t, err)

	client := &CraneDockerClient{
		sharedHttpClient:         httpClient,
		swarmManagerHttpEndpoint: server.URL,
	}

	_, err = client.CreateService(swarm.ServiceSpec{}, types.ServiceCreateOptions{})
	assert.Nil(t, err)

	client = &CraneDockerClient{
		sharedHttpClient:         httpClient,
		swarmManagerHttpEndpoint: "errorurl",
	}

	_, err = client.CreateService(swarm.ServiceSpec{}, types.ServiceCreateOptions{})
	assert.NotNil(t, err)

	server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`error response`))
	}))
	client = &CraneDockerClient{
		sharedHttpClient:         httpClient,
		swarmManagerHttpEndpoint: server.URL,
	}

	_, err = client.CreateService(swarm.ServiceSpec{}, types.ServiceCreateOptions{})
	assert.NotNil(t, err)
}

func TestListServiceSpec(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`[{"ID":"e90302"}]`))
	}))
	defer server.Close()

	httpClient, err := NewHttpClient()
	assert.Nil(t, err)
	client := &CraneDockerClient{
		sharedHttpClient:         httpClient,
		swarmManagerHttpEndpoint: server.URL,
	}
	_, err = client.ListServiceSpec(types.ServiceListOptions{})
	assert.Nil(t, err)

	client = &CraneDockerClient{
		sharedHttpClient:         httpClient,
		swarmManagerHttpEndpoint: "error url",
	}
	_, err = client.ListServiceSpec(types.ServiceListOptions{})
	assert.NotNil(t, err)

	server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`error response`))
	}))
	client = &CraneDockerClient{
		sharedHttpClient:         httpClient,
		swarmManagerHttpEndpoint: server.URL,
	}
	_, err = client.ListServiceSpec(types.ServiceListOptions{})
	assert.NotNil(t, err)
}

func TestListService(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`[{"ID":"e90302"}]`))
	}))
	defer server.Close()

	httpClient, err := NewHttpClient()
	assert.Nil(t, err)
	client := &CraneDockerClient{
		sharedHttpClient:         httpClient,
		swarmManagerHttpEndpoint: server.URL,
	}
	_, err = client.ListService(types.ServiceListOptions{})
	assert.Nil(t, err)

	server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`error response`))
	}))
	client = &CraneDockerClient{
		sharedHttpClient:         httpClient,
		swarmManagerHttpEndpoint: "testurl",
	}
	_, err = client.ListService(types.ServiceListOptions{})
	assert.NotNil(t, err)
}

func FakeErrorListNodes(ctx *gin.Context) {
	body := ``
	ctx.JSON(http.StatusOK, body)
}

func FakeListNodes(ctx *gin.Context) {
	body := `
	[
		{
    	    "ID":"1t6jojzasio4veexyubvic4j2",
    	    "Version":{
    	        "Index":26607
    	    },
    	    "CreatedAt":"2016-08-26T08:00:24.466491891Z",
    	    "UpdatedAt":"2016-09-08T05:23:49.697933079Z",
    	    "Spec":{
    	        "Labels":{
    	            "dm.reserved.node.endpoint":"http://192.168.59.103:2376"
    	        },
    	        "Role":"worker",
    	        "Availability":"active"
    	    },
    	    "Description":{
    	        "Hostname":"192.168.59.013",
    	        "Platform":{
    	            "Architecture":"x86_64",
    	            "OS":"linux"
    	        },
    	        "Resources":{
    	            "NanoCPUs":2000000000,
    	            "MemoryBytes":3975561216
    	        },
    	        "Engine":{
    	            "EngineVersion":"1.12.0",
    	            "Plugins":[
    	                {
    	                    "Type":"Network",
    	                    "Name":"bridge"
    	                },
    	                {
    	                    "Type":"Network",
    	                    "Name":"host"
    	                },
    	                {
    	                    "Type":"Network",
    	                    "Name":"null"
    	                },
    	                {
    	                    "Type":"Network",
    	                    "Name":"overlay"
    	                },
    	                {
    	                    "Type":"Volume",
    	                    "Name":"local"
    	                }
    	            ]
    	        }
    	    },
    	    "Status":{
    	        "State":"ready"
    	    }
    	},
    	{
    	    "ID":"dbspw1g0sjee8ja1khx2w0xtt",
    	    "Version":{
    	        "Index":26603
    	    },
    	    "CreatedAt":"2016-08-26T07:59:50.685235915Z",
    	    "UpdatedAt":"2016-09-08T05:23:36.061728082Z",
    	    "Spec":{
    	        "Labels":{
    	            "dm.reserved.node.endpoint":"192.168.59.104:2376"
    	        },
    	        "Role":"manager",
    	        "Availability":"active"
    	    },
    	    "Description":{
    	        "Hostname":"localhost",
    	        "Platform":{
    	            "Architecture":"x86_64",
    	            "OS":"linux"
    	        },
    	        "Resources":{
    	            "NanoCPUs":2000000000,
    	            "MemoryBytes":3975561216
    	        },
    	        "Engine":{
    	            "EngineVersion":"1.12.0",
    	            "Plugins":[
    	                {
    	                    "Type":"Network",
    	                    "Name":"bridge"
    	                },
    	                {
    	                    "Type":"Network",
    	                    "Name":"host"
    	                },
    	                {
    	                    "Type":"Network",
    	                    "Name":"null"
    	                },
    	                {
    	                    "Type":"Network",
    	                    "Name":"overlay"
    	                },
    	                {
    	                    "Type":"Volume",
    	                    "Name":"local"
    	                }
    	            ]
    	        }
    	    },
    	    "Status":{
    	        "State":"ready"
    	    },
    	    "ManagerStatus":{
    	        "Leader":true,
    	        "Reachability":"reachable",
    	        "Addr":"192.168.59.104:2377"
    	    }
    	}
	]
	`
	ctx.Writer.Header().Set("Content-Type", "application/json")
	ctx.Writer.Write([]byte(body))
}

func FakeErrorListTasks(ctx *gin.Context) {
	body := ``
	ctx.JSON(http.StatusOK, body)
}

func FakeListTasks(ctx *gin.Context) {
	body := `
[
  {
    "ID":"9s978twp81zo0f34r4bdwz0pg",
    "Version":{
      "Index":51
    },
    "CreatedAt":"2016-09-12T15:16:56.784488228Z",
    "UpdatedAt":"2016-09-13T06:01:24.786916351Z",
    "Spec":{
      "ContainerSpec":{
        "Image":"blackicebird/2048"
      }
    },
    "ServiceID":"1ns52ng0jqs7c3lw8ydxbrga4",
    "Slot":1,
    "NodeID":"28u7wv664h5rrlw6jfs51pi8o",
    "Status":{
      "Timestamp":"2016-09-13T06:01:24.720077916Z",
      "State":"running",
      "Message":"started",
      "ContainerStatus":{
        "ContainerID":"6f0d3fab83b1526d71fb1fa10e465d3b585b0d0971b92d0cdbba46e2a78e7a7d",
        "PID":12574
      }
    },
    "DesiredState":"running",
    "NetworksAttachments":[
      {
        "Network":{
          "ID":"c0it8e2mhwcnbebm494639496",
          "Version":{
            "Index":41
          },
          "CreatedAt":"2016-09-12T03:41:21.32445873Z",
          "UpdatedAt":"2016-09-13T06:01:17.602780434Z",
          "Spec":{
            "Name":"ingress",
            "Labels":{
              "com.docker.swarm.internal":"true"
            },
            "DriverConfiguration":{

            },
            "IPAMOptions":{
              "Driver":{

              },
              "Configs":[
                {
                  "Subnet":"10.255.0.0/16",
                  "Gateway":"10.255.0.1"
                }
              ]
            }
          },
          "DriverState":{
            "Name":"overlay",
            "Options":{
              "com.docker.network.driver.overlay.vxlanid_list":"256"
            }
          },
          "IPAMOptions":{
            "Driver":{
              "Name":"default"
            },
            "Configs":[
              {
                "Subnet":"10.255.0.0/16",
                "Gateway":"10.255.0.1"
              }
            ]
          }
        },
        "Addresses":[
          "10.255.0.6/16"
        ]
      },
      {
        "Network":{
          "ID":"c0it8e2mhwcnbebm494639496",
          "Version":{
            "Index":41
          },
          "CreatedAt":"2016-09-12T03:41:21.32445873Z",
          "UpdatedAt":"2016-09-13T06:01:17.602780434Z",
          "Spec":{
            "Name":"ingress",
            "Labels":{
              "com.docker.swarm.internal":"true"
            },
            "DriverConfiguration":{

            },
            "IPAMOptions":{
              "Driver":{

              },
              "Configs":[
                {
                  "Subnet":"10.255.0.0/16",
                  "Gateway":"10.255.0.1"
                }
              ]
            }
          },
          "DriverState":{
            "Name":"overlay",
            "Options":{
              "com.docker.network.driver.overlay.vxlanid_list":"256"
            }
          },
          "IPAMOptions":{
            "Driver":{
              "Name":"default"
            },
            "Configs":[
              {
                "Subnet":"10.255.0.0/16",
                "Gateway":"10.255.0.1"
              }
            ]
          }
        },
        "Addresses":[
          "10.255.0.7/16"
        ]
      }
    ]
  },
  {
    "ID":"b6pv8tjozbpiqj2xnjj6v9ogx",
    "Version":{
      "Index":48
    },
    "CreatedAt":"2016-09-12T06:18:17.046935652Z",
    "UpdatedAt":"2016-09-13T06:01:18.72012925Z",
    "Spec":{
      "ContainerSpec":{
        "Image":"blackicebird/2048"
      }
    },
    "ServiceID":"1ns52ng0jqs7c3lw8ydxbrga4",
    "Slot":1,
    "NodeID":"28u7wv664h5rrlw6jfs51pi8o",
    "Status":{
      "Timestamp":"2016-09-13T06:01:13.593525606Z",
      "State":"running",
      "Message":"started",
      "Err":"task: non-zero exit (1)",
      "ContainerStatus":{
        "ContainerID":"419e547b71a69801e5217593fb4da26616552859526890564bb0844d5ca1e00f",
        "ExitCode":1
      }
    },
    "DesiredState":"shutdown",
    "NetworksAttachments":[
      {
        "Network":{
          "ID":"c0it8e2mhwcnbebm494639496",
          "Version":{
            "Index":41
          },
          "CreatedAt":"2016-09-12T03:41:21.32445873Z",
          "UpdatedAt":"2016-09-13T06:01:17.602780434Z",
          "Spec":{
            "Name":"ingress",
            "Labels":{
              "com.docker.swarm.internal":"true"
            },
            "DriverConfiguration":{

            },
            "IPAMOptions":{
              "Driver":{

              },
              "Configs":[
                {
                  "Subnet":"10.255.0.0/16",
                  "Gateway":"10.255.0.1"
                }
              ]
            }
          },
          "DriverState":{
            "Name":"overlay",
            "Options":{
              "com.docker.network.driver.overlay.vxlanid_list":"256"
            }
          },
          "IPAMOptions":{
            "Driver":{
              "Name":"default"
            },
            "Configs":[
              {
                "Subnet":"10.255.0.0/16",
                "Gateway":"10.255.0.1"
              }
            ]
          }
        },
        "Addresses":[
          "10.255.0.4/16"
        ]
      },
      {
        "Network":{
          "ID":"c0it8e2mhwcnbebm494639496",
          "Version":{
            "Index":41
          },
          "CreatedAt":"2016-09-12T03:41:21.32445873Z",
          "UpdatedAt":"2016-09-13T06:01:17.602780434Z",
          "Spec":{
            "Name":"ingress",
            "Labels":{
              "com.docker.swarm.internal":"true"
            },
            "DriverConfiguration":{

            },
            "IPAMOptions":{
              "Driver":{

              },
              "Configs":[
                {
                  "Subnet":"10.255.0.0/16",
                  "Gateway":"10.255.0.1"
                }
              ]
            }
          },
          "DriverState":{
            "Name":"overlay",
            "Options":{
              "com.docker.network.driver.overlay.vxlanid_list":"256"
            }
          },
          "IPAMOptions":{
            "Driver":{
              "Name":"default"
            },
            "Configs":[
              {
                "Subnet":"10.255.0.0/16",
                "Gateway":"10.255.0.1"
              }
            ]
          }
        },
        "Addresses":[
          "10.255.0.5/16"
        ]
      }
    ]
  }
]
`
	ctx.Writer.Header().Set("Content-Type", "application/json")
	ctx.Writer.Write([]byte(body))
}

func TestGetServiceStatus(t *testing.T) {
	var services []swarm.Service
	var RqServiceSt, RqServiceSt2 []ServiceStatus

	serviceContent := `
	[
	    {
		"ID":"1ns52ng0jqs7c3lw8ydxbrga4",
		"Version":{
		    "Index":43
		},
		"CreatedAt":"2016-09-12T06:18:17.03174955Z",
		"UpdatedAt":"2016-09-13T06:01:17.607771366Z",
		"Spec":{
		    "Name":"test_2048",
		    "Labels":{
			"com.docker.stack.namespace":"test",
			"crane.reserved.permissions.1.r":"true",
			"crane.reserved.permissions.1.w":"true",
			"crane.reserved.permissions.1.x":"true",
			"name":"2048"
		    },
		    "TaskTemplate":{
			"ContainerSpec":{
			    "Image":"blackicebird/2048"
			}
		    },
		    "Mode":{
			"Replicated":{
			    "Replicas":1
			}
		    },
		    "Networks":[
			{
			    "Target":"c0it8e2mhwcnbebm494639496",
			    "Aliases":[
				"2048"
			    ]
			}
		    ],
		    "EndpointSpec":{
			"Mode":"vip",
			"Ports":[
			    {
				"Name":"pbport",
				"Protocol":"tcp",
				"TargetPort":80,
				"PublishedPort":8000
			    }
			]
		    }
		},
		"Endpoint":{
		    "Spec":{
			"Mode":"vip",
			"Ports":[
			    {
				"Name":"pbport",
				"Protocol":"tcp",
				"TargetPort":80,
				"PublishedPort":8000
			    }
			]
		    },
		    "Ports":[
			{
			    "Name":"pbport",
			    "Protocol":"tcp",
			    "TargetPort":80,
			    "PublishedPort":8000
			}
		    ],
		    "VirtualIPs":[
			{
			    "NetworkID":"c0it8e2mhwcnbebm494639496",
			    "Addr":"10.255.0.2/16"
			}
		    ]
		},
		"UpdateStatus":{
		    "StartedAt":"0001-01-01T00:00:00Z",
		    "CompletedAt":"0001-01-01T00:00:00Z"
		}
	    }
	]
	`
	json.Unmarshal([]byte(serviceContent), &services)

	// test error ListNodes
	httpClient, err := NewHttpClient()
	assert.Nil(t, err)

	router := gin.New()
	router.GET("/nodes", FakeErrorListNodes)
	router.GET("/tasks", FakeListTasks)

	server := httptest.NewServer(router)

	client := &CraneDockerClient{
		sharedHttpClient:         httpClient,
		swarmManagerHttpEndpoint: server.URL,
	}
	_, err = client.GetServicesStatus(services)
	assert.NotNil(t, err)
	server.Close()

	// test error ListTasks
	router = gin.New()
	router.GET("/nodes", FakeListNodes)
	router.GET("/tasks", FakeErrorListTasks)

	server = httptest.NewServer(router)

	client = &CraneDockerClient{
		sharedHttpClient:         httpClient,
		swarmManagerHttpEndpoint: server.URL,
	}
	_, err = client.GetServicesStatus(services)
	assert.NotNil(t, err)
	server.Close()

	// test ok
	router = gin.New()
	router.GET("/nodes", FakeListNodes)
	router.GET("/tasks", FakeListTasks)

	server = httptest.NewServer(router)

	client = &CraneDockerClient{
		sharedHttpClient:         httpClient,
		swarmManagerHttpEndpoint: server.URL,
	}
	servicesSt, err := client.GetServicesStatus(services)
	assert.Nil(t, err)

	RqServiceStBody := `
[
    {
        "ID":"1ns52ng0jqs7c3lw8ydxbrga4",
        "Name":"test_2048",
        "NumTasksRunning":0,
        "NumTasksTotal":1,
        "Image":"blackicebird/2048",
        "Command":"",
        "CreatedAt":"2016-09-12T06:18:17.03174955Z",
        "UpdatedAt":"2016-09-13T06:01:17.607771366Z",
        "LimitCpus":0,
        "LimitMems":0,
        "ReserveCpus":0,
        "ReserveMems":0,
        "IPs":[
            "",
            ""
        ],
        "Ports":[
            8000
        ]
    }
]
`
	// test TaskTemplate.Resources
	json.Unmarshal([]byte(RqServiceStBody), &RqServiceSt)

	assert.Equal(t, servicesSt, RqServiceSt)
	server.Close()

	serviceContent = `
	[
	    {
		"ID":"1ns52ng0jqs7c3lw8ydxbrga4",
		"Version":{
		    "Index":43
		},
		"CreatedAt":"2016-09-12T06:18:17.03174955Z",
		"UpdatedAt":"2016-09-13T06:01:17.607771366Z",
		"Spec":{
		    "Name":"test_2048",
		    "Labels":{
			"com.docker.stack.namespace":"test",
			"crane.reserved.permissions.1.r":"true",
			"crane.reserved.permissions.1.w":"true",
			"crane.reserved.permissions.1.x":"true",
			"name":"2048"
		    },
		    "TaskTemplate":{
			"ContainerSpec":{
			    "Image":"blackicebird/2048"
			},
			"Resources":{
				"Limits":{
					"NanoCPUs":1,
					"MemoryBytes":1024
				},
				"Reservations":{
					"NanoCPUs":1,
					"MemoryBytes":1024
				}
			}
		    },
		    "Mode":{
			"Global":{
			}
		    },
		    "Networks":[
			{
			    "Target":"c0it8e2mhwcnbebm494639496",
			    "Aliases":[
				"2048"
			    ]
			}
		    ],
		    "EndpointSpec":{
			"Mode":"vip",
			"Ports":[
			    {
				"Name":"pbport",
				"Protocol":"tcp",
				"TargetPort":80,
				"PublishedPort":8000
			    }
			]
		    }
		},
		"Endpoint":{
		    "Spec":{
			"Mode":"vip",
			"Ports":[
			    {
				"Name":"pbport",
				"Protocol":"tcp",
				"TargetPort":80,
				"PublishedPort":8000
			    }
			]
		    },
		    "Ports":[
			{
			    "Name":"pbport",
			    "Protocol":"tcp",
			    "TargetPort":80,
			    "PublishedPort":8000
			}
		    ],
		    "VirtualIPs":[
			{
			    "NetworkID":"c0it8e2mhwcnbebm494639496",
			    "Addr":"10.255.0.2/16"
			}
		    ]
		},
		"UpdateStatus":{
		    "StartedAt":"0001-01-01T00:00:00Z",
		    "CompletedAt":"0001-01-01T00:00:00Z"
		}
	    }
	]
	`
	json.Unmarshal([]byte(serviceContent), &services)

	router = gin.New()
	router.GET("/nodes", FakeListNodes)
	router.GET("/tasks", FakeListTasks)

	server = httptest.NewServer(router)

	client = &CraneDockerClient{
		sharedHttpClient:         httpClient,
		swarmManagerHttpEndpoint: server.URL,
	}
	servicesSt2, err := client.GetServicesStatus(services)
	assert.Nil(t, err)

	RqServiceStBody2 := `
[
    {
        "ID":"1ns52ng0jqs7c3lw8ydxbrga4",
        "Name":"test_2048",
        "NumTasksRunning":0,
        "NumTasksTotal":1,
        "Image":"blackicebird/2048",
        "Command":"",
        "CreatedAt":"2016-09-12T06:18:17.03174955Z",
        "UpdatedAt":"2016-09-13T06:01:17.607771366Z",
        "LimitCpus":0,
        "LimitMems":0,
        "ReserveCpus":1,
        "ReserveMems":1024,
        "IPs":[
            "",
            ""
        ],
        "Ports":[
            8000
        ]
    }
]
`

	json.Unmarshal([]byte(RqServiceStBody2), &RqServiceSt2)

	assert.Equal(t, servicesSt2, RqServiceSt2)
	server.Close()
}
