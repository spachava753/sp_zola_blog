+++
title = "QuickFixJ and Spring Boot tutorial"
date = 2018-11-17
[taxonomies]
tags = ["FIX", "Java", "Spring boot", "Finance", "Tutorial"]
+++

# What is QuickFixJ

## What is a FIX

FIX stands for Financial Information exchange, and the protocol was put forth in 1992 for real-time exchange of information that is related to the securities transactions and markets. There are multiple version of this protocol, and in each there are specific fields that are required. The protocol exchanges a myriad of types of messages, from orders to financial news.

## What is a fix engine

A fix engine is what is used for actually communicating across the FIX protocol. QuickFix is an example of a fix engine. QuickFixJ is the Java variant of Quickfix, which was originally written in C++. Some institutions write their own proprietary software, for robust performance, but we will be using QuickFixJ with Spring Boot for rapid development. There are two types of fix engines: clients and servers. In QuickFixJ, initiators are the clients that connect to acceptors, and acceptors process the message sent by the initiator, such as executing a new order. Most of the time, we won't need to worry about acceptors as we will be the ones connecting to them, so this tutorial will focus on writing an initiator, although most of the code will still work if you want to write an acceptor with QuickFixJ.

## QuickFixJ notes

We'll be using the latest QuickFixJ version, but many institutions use older versions so in terms of reliability, older version of QuickFixJ are battle-tested and can endure the river of time. This tutorial is only a basic article on integrating QuickFixJ into Spring Boot, not an extensive analysis and instruction set on building a production-ready OMS (Order Management Systems).

# What we'll be using

- An IDE, I used IntelliJ
- Maven (You can use Gradle if you prefer)
- Spring Boot 2.1.0
- QuickFixJ 2.1.0
- Project Lombok 1.18.4

# Creating your own Fix Engine

## Creating the project

### Creating a Java project

A popular IDE such as IntelliJ should allow you to create a new Maven Project, which will initialize the directory structure and create the pom.xml file for you. In IntelliJ, its File -> New -> Select Maven Project on the left and follow the steps.

### Creating the POM

```XML
<?xml version="1.0" encoding="UTF-8"?>
<project xmlns="http://maven.apache.org/POM/4.0.0" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
    xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">
    <modelVersion>4.0.0</modelVersion>

    <groupId>com.example</groupId>
    <artifactId>QuickFixJ-spring-boot</artifactId>
    <version>0.0.1-SNAPSHOT</version>
    <packaging>jar</packaging>

    <name>QuickFixJ-spring-boot-example</name>
    <description>A demo project with QuickFixJ and Spring Boot to demonstrate how they integrate together.</description>

    <parent>
        <groupId>org.springframework.boot</groupId>
        <artifactId>spring-boot-starter-parent</artifactId>
        <version>2.1.0.RELEASE</version>
        <relativePath/> <!-- lookup parent from repository -->
    </parent>

    <properties>
        <project.build.sourceEncoding>UTF-8</project.build.sourceEncoding>
        <project.reporting.outputEncoding>UTF-8</project.reporting.outputEncoding>
        <java.version>1.8</java.version>
    </properties>

    <dependencies>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-web</artifactId>
        </dependency>
        <dependency>
            <groupId>org.QuickFixJ</groupId>
            <artifactId>QuickFixJ-core</artifactId>
            <version>2.1.0</version>
        </dependency>
        <dependency>
            <groupId>org.QuickFixJ</groupId>
            <artifactId>QuickFixJ-messages-all</artifactId>
            <version>2.1.0</version>
        </dependency>
        <dependency>
            <groupId>org.projectlombok</groupId>
            <artifactId>lombok</artifactId>
            <version>1.18.4</version>
            <scope>provided</scope>
        </dependency>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-devtools</artifactId>
            <scope>runtime</scope>
        </dependency>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-test</artifactId>
            <scope>test</scope>
        </dependency>
    </dependencies>

    <build>
        <plugins>
            <plugin>
                <groupId>org.springframework.boot</groupId>
                <artifactId>spring-boot-maven-plugin</artifactId>
            </plugin>
        </plugins>
    </build>

</project>
```

Your pom.xml should look something like this, except for the groupId and the artifactId. Spring Boot Devtools is optional, and you can remove it if you want. However, it is convenient for easy restart of our application. All you have to do is build the application in an IDE. If you encounter problems, then remove it. This is all we will need for creating our app.

### Making the Spring Boot app

We will first write out our main method first and annotate it with **@SpringBootApplication**.

```Java
package com.example.QuickFixJspringboot;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
public class QuickFixJSpringBootExampleApplication {

    public static void main(String[] args) {
        SpringApplication.run(QuickFixJSpringBootExampleApplication.class, args);
    }
}
```

Run this with **mvn clean spring-boot:run** and the application should start on port 8080. If you run into problems, check that nothing is running in port 8080 first, so there are no conflicts. The application doesn't do anything, so we can't really see the results of our work yet. Let's add a RestController.

```Java
package com.example.QuickFixJspringboot.web.rest;

import lombok.extern.slf4j.Slf4j;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@Slf4j
public class StatusController {

    @GetMapping("/status")
    public ResponseEntity<String> status() {
        log.info("-----Call localhost:8080/status-----");
        return new ResponseEntity<>("Up and Running!", HttpStatus.OK);
    }

}
```

The **@RestController** annotation allows Spring Boot to recognize this class as a rest controller, and the **@GetMapping** annotation is like the **@RequestMapping** annotation, but it specifies that only GET requests are accepted. The **@Slf4j** annotation is part of [Project Lombok](https://projectlombok.org/), a really cool project that allows use to circumvent writing a lot of boilerplate code like loggers, getters, setters, etc. This specific annotation atomatically creates an Slf4J logger for us to use. Make sure to find out if your IDE supports Lombok. For IntelliJ, there is a free Lombok plugin. Now try running **mvn clean spring-boot:run** again, and go to [localhost:8080/status](http://localhost:8080/status), and you'll find the status of the application. Now its time for us to integrate QuickFixJ into the application.

## Integrating QuickFixJ into our application

QuickFixJ allows us to create our fix engine without much boilerplate code. In fact, with only one class and a main method, our fix engine would be up and running. In order to run, QuickFixJ utilizes Java sockets, so all we need to say **start()**. However, the process becomes more complex when you try to integrate with Spring Boot. Spring Boot force closes the sockets, which means that our fix engine is no longer in session. Therefore, we need to capitalize on Spring Boot's lifecycle events and annotations to make Spring Boot recognize our QuickFixJ classes as a valid Java bean.

### Creating the FixApplication

```Java
package com.example.QuickFixJspringboot.fix;

import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Component;
import quickfix.*;

@Slf4j
@Component
public class ClientFixApp implements Application {

    @Override
    public void onCreate(SessionID sessionID) {
        log.info("--------- onCreate ---------");
    }

    @Override
    public void onLogon(SessionID sessionID) {
        log.info("--------- onLogon ---------");
    }

    @Override
    public void onLogout(SessionID sessionID) {
        log.info("--------- onLogout ---------");
    }

    @Override
    public void toAdmin(Message message, SessionID sessionID) {
        log.info("--------- toAdmin ---------");
    }

    @Override
    public void fromAdmin(Message message, SessionID sessionID) throws FieldNotFound, IncorrectDataFormat, IncorrectTagValue, RejectLogon {
        log.info("--------- fromAdmin ---------");
    }

    @Override
    public void toApp(Message message, SessionID sessionID) throws DoNotSend {
        log.info("--------- toApp ---------");
    }

    @Override
    public void fromApp(Message message, SessionID sessionID) throws FieldNotFound, IncorrectDataFormat, IncorrectTagValue, UnsupportedMessageType {
        log.info("--------- fromApp ---------");
    }
}

```

These methods are what is used to do specific steps for each event in the FIX protocol. Look forward to a future tutorial where I will take the Banzai client given to us by QuickFixJ and convert it to be used with Spring Boot.

### Creating the Initiator configuration

```Java
package com.example.QuickFixJspringboot.config;

import com.example.QuickFixJspringboot.fix.ClientFixApp;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import quickfix.*;

import java.io.FileInputStream;
import java.io.FileNotFoundException;

@Configuration
@Slf4j
public class FixConfig {

    private final String fileName = "initiator.cfg";
    @Autowired
    private ClientFixApp application;

    @Bean
    public ThreadedSocketInitiator threadedSocketInitiator(){
        ThreadedSocketInitiator threadedSocketInitiator = null;

        try {
            SessionSettings settings = new SessionSettings(new FileInputStream(fileName));
            MessageStoreFactory storeFactory = new FileStoreFactory(settings);
            LogFactory logFactory = new FileLogFactory(settings);
            MessageFactory messageFactory = new DefaultMessageFactory();
            threadedSocketInitiator = new ThreadedSocketInitiator(application, storeFactory, settings, logFactory, messageFactory);
        } catch (ConfigError configError) {
            configError.printStackTrace();
        } catch (FileNotFoundException e) {
            e.printStackTrace();
        }

        return threadedSocketInitiator;
    }
}
```

This is the important part. This class tells Spring Boot how to create our ThreadedSocketInitiator bean, so we can **@Autowire** it later. The **@Bean** annotation is what defines our custom initialization of the ThreadedSocketInitator, and the **@Configuration** annotation allows Spring Boot to pick up that this class is important during the initalization of the app, and should be used to configure specfic user-defined beans.

```
[default]
FileStorePath=target/data/banzai
FileLogPath=target/data/banzai
ConnectionType=initiator
SenderCompID=CLIENTFIX
TargetCompID=EXEC
SocketConnectHost=localhost
StartTime=00:00:00
EndTime=00:00:00
HeartBtInt=30
ReconnectInterval=5

[session]
BeginString=FIX.4.0
SocketConnectPort=9876

[session]
BeginString=FIX.4.1
SocketConnectPort=9877

[session]
BeginString=FIX.4.2
SocketConnectPort=9878

[session]
BeginString=FIX.4.3
SocketConnectPort=9879

[session]
BeginString=FIX.4.4
SocketConnectPort=9880

[session]
BeginString=FIXT.1.1
DefaultApplVerID=FIX.5.0
SocketConnectPort=9881
```

This is the config file our application uses. Include this in the root folder of the project and call this **initiator.cfg**. This config was taken directly from the provided Banzai client in the QuickFixJ examples. This will connect to the example Executor that QuickFixJ implemented, so you can connect them and test the connection.

```java
package com.example.QuickFixJspringboot.listener;

import com.example.QuickFixJspringboot.fix.ClientFixApp;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.ApplicationListener;
import org.springframework.context.event.ContextRefreshedEvent;
import org.springframework.context.event.ContextStartedEvent;
import org.springframework.scheduling.annotation.Scheduled;
import org.springframework.stereotype.Component;
import quickfix.ConfigError;
import quickfix.Session;
import quickfix.SessionID;
import quickfix.ThreadedSocketInitiator;

import java.time.LocalDateTime;
import java.time.format.DateTimeFormatter;

@Component
@Slf4j
public class AppLogonListener implements ApplicationListener<ContextRefreshedEvent> {

    @Autowired
    private ThreadedSocketInitiator threadedSocketInitiator;

    @Autowired
    private ClientFixApp application;

    private boolean initiatorStarted = false;

    @Override
    public void onApplicationEvent(ContextRefreshedEvent refreshedEvent) {
        startFixInitiator();
    }

    private void startFixInitiator (){
        if(!initiatorStarted) {
            try {
                threadedSocketInitiator.start();
                log.info("--------- ThreadedSocketInitiator started ---------");
                initiatorStarted = true;
            } catch (ConfigError configError) {
                configError.printStackTrace();
                log.error("--------- ThreadedSocketInitiator ran into an error ---------");
            }
        } else {
            logon();
        }
    }

    private void logon (){
        if(threadedSocketInitiator.getSessions() != null && threadedSocketInitiator.getSessions().size() > 0) {
            for (SessionID sessionID: threadedSocketInitiator.getSessions()) {
                Session.lookupSession(sessionID).logon();
            }
            log.info("--------- ThreadedSocketInitiator logged on to sessions. Size: " + threadedSocketInitiator.getSessions().size() + " ---------");
        }
    }

    @Scheduled(fixedRate = 5000)
    public void clientStatus(){
        log.info("Client Status | Logged on: {}. Current Time: {}", threadedSocketInitiator.isLoggedOn(), LocalDateTime.now().format(DateTimeFormatter.ofPattern("yyyy-MM-dd HH:mm:ss.SSS")));
    }
}
```

This is the last part of this tutorial, the logic that starts the fix engine. The **@Component** annotation is similar to the **@Configuration** annotation in that Spring Boot will recognize this class as special. The class implements the `ApplicationListener<ContextRefreshedEvent>` interface to hook on to the Spring lifecycle. The method **onApplicationEvent(ContextRefreshedEvent refreshedEvent)** will execute when the ApplicationContext refreshes, Spring Boot calls any listeners on the **ContextRefreshedEvent** flag. We also add a method that reports the status of our initiator every five seconds, under the **@Scheduled** annotation.

# Conclusion

So far we have created our Spring Boot app, and if you test it against the Executor that QuickFixJ implemented, you'll find that the app connects fine. I would like to point out that this tutorial showed how to initalize your classes with a custom bean configuration, as well as how to make Spring Boot recognize and **not** shutdown our Java sockets. I hope you learned how to not only integrate QuickFixJ into Spring Boot, but also about Project Lombok and various annotation of Spring Boot. Unfortunately, we didn't implement any functionality for **ClientFixApp**, so it doesn't do anything yet. In future posts, I will be posting a step by step tutorial on writing a Spring Boot Banzai client for the Executor.

### [Github link](https://github.com/spachava753/quickfixj-spring-boot)