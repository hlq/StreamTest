package org.stream.cloud.eurekatwoclient;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.cloud.context.config.annotation.RefreshScope;
import org.springframework.cloud.netflix.eureka.EnableEurekaClient;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

//@RefreshScope
@SpringBootApplication
@EnableEurekaClient
@RestController
@RefreshScope
public class EurekaTwoClientApplication {

    public static void main(String[] args) {
        SpringApplication.run(EurekaTwoClientApplication.class, args);
    }

    @Value("${server.port}")
    String port;

    @Value("${to}")
    private String to;

    @RequestMapping("/hi")
    public String home(@RequestParam(value = "name", defaultValue = "Stream") String name) {
        return "hi " + name + " ,i am from port:" + port + " !!! " + to;
    }

}
