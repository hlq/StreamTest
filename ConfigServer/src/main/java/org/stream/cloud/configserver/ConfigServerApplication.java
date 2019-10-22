package org.stream.cloud.configserver;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.cloud.client.discovery.EnableDiscoveryClient;
import org.springframework.cloud.config.server.EnableConfigServer;
import org.springframework.cloud.netflix.eureka.EnableEurekaClient;

@SpringBootApplication
@EnableConfigServer
@EnableEurekaClient
@EnableDiscoveryClient
public class ConfigServerApplication {
    /**
     * http://localhost:5001/service-hi/dev/master git 内容  
     * http://localhost:5001/actuator/bus-refresh post 更新config
     * 
     * 
     */
    
    
    
    public static void main(String[] args) {
        SpringApplication.run(ConfigServerApplication.class, args);
    }

}


