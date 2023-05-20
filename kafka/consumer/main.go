package main

import (
	"crypto/tls"
	"fmt"
	"github.com/Shopify/sarama"
	"golang.org/x/net/proxy"
	"net"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	config := sarama.Config{
		Admin: struct {
			Retry struct {
				Max     int
				Backoff time.Duration
			}
			Timeout time.Duration
		}{},
		Net: struct {
			MaxOpenRequests int
			DialTimeout     time.Duration
			ReadTimeout     time.Duration
			WriteTimeout    time.Duration
			TLS             struct {
				Enable bool
				Config *tls.Config
			}
			SASL struct {
				Enable                   bool
				Mechanism                sarama.SASLMechanism
				Version                  int16
				Handshake                bool
				AuthIdentity             string
				User                     string
				Password                 string
				SCRAMAuthzID             string
				SCRAMClientGeneratorFunc func() sarama.SCRAMClient
				TokenProvider            sarama.AccessTokenProvider
				GSSAPI                   sarama.GSSAPIConfig
			}
			KeepAlive time.Duration
			LocalAddr net.Addr
			Proxy     struct {
				Enable bool
				Dialer proxy.Dialer
			}
		}{},
		Metadata: struct {
			Retry struct {
				Max         int
				Backoff     time.Duration
				BackoffFunc func(retries int, maxRetries int) time.Duration
			}
			RefreshFrequency       time.Duration
			Full                   bool
			Timeout                time.Duration
			AllowAutoTopicCreation bool
		}{},
		Consumer: struct {
			Group struct {
				Session struct {
					Timeout time.Duration // 死亡时间
				}
				Heartbeat struct {
					Interval time.Duration //心跳间隔
				}
				Rebalance struct {
					Strategy        sarama.BalanceStrategy
					GroupStrategies []sarama.BalanceStrategy
					Timeout         time.Duration
					Retry           struct {
						Max     int
						Backoff time.Duration
					}
				}
				Member struct {
					UserData []byte
				}
				InstanceId          string
				ResetInvalidOffsets bool
			}
			Retry struct {
				Backoff     time.Duration
				BackoffFunc func(retries int) time.Duration
			}
			Fetch struct {
				Min     int32 // 每次拉去的最小字节数
				Default int32
				Max     int32 // 每次拉去的最长时间
			}
			MaxWaitTime       time.Duration //最大等待时间
			MaxProcessingTime time.Duration
			Return            struct {
				Errors bool
			}
			Offsets struct {
				CommitInterval time.Duration
				AutoCommit     struct {
					Enable   bool          // 是否自动提交
					Interval time.Duration //自动提交的频率
				}
				Initial   int64
				Retention time.Duration
				Retry     struct {
					Max int
				}
			}
			IsolationLevel sarama.IsolationLevel
			Interceptors   []sarama.ConsumerInterceptor
		}{},
		ClientID:           "",
		RackID:             "",
		ChannelBufferSize:  0,
		ApiVersionsRequest: false,
		Version:            sarama.KafkaVersion{},
		MetricRegistry:     nil,
	}

	consumer, err := sarama.NewConsumer([]string{"118.126.89.12:9092"}, &config)
	if err != nil {
		fmt.Println("consumer connect err:", err)
		return
	}
	defer consumer.Close()

	//获取 kafka 主题
	partitions, err := consumer.Partitions("cctv1")
	if err != nil {
		fmt.Println("get partitions failed, err:", err)
		return
	}

	for _, p := range partitions {
		//sarama.OffsetNewest：从当前的偏移量开始消费，sarama.OffsetOldest：从最老的偏移量开始消费
		partitionConsumer, err := consumer.ConsumePartition("cctv1", p, sarama.OffsetNewest)
		if err != nil {
			fmt.Println("partitionConsumer err:", err)
			continue
		}
		wg.Add(1)
		go func() {
			for m := range partitionConsumer.Messages() {
				fmt.Printf("key: %s, text: %s, offset: %d\n", string(m.Key), string(m.Value), m.Offset)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
