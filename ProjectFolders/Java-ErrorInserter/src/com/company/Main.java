package com.company;

import com.rabbitmq.client.*;
import com.rabbitmq.client.Connection;

import java.io.IOException;
import java.util.Map;
import java.sql.*;

public class Main {

    private static final String TASK_QUE_NAME = "HNError";
    private static String DBUser;
    private static String DBPass;
    private static String DBIP;
    private static String RabbitUser;
    private static String RabbitPass;
    private static String RabbitIP;




    public static void main(String[] args) throws Exception {

        DBUser = args[0];
        DBPass = args[1];
        DBIP = args[2];
        RabbitUser = args[3];
        RabbitPass = args[4];
        RabbitIP = args[5];


        ConnectionFactory factory = new ConnectionFactory();
        factory.setHost(RabbitIP);
        factory.setUsername(RabbitUser);
        factory.setPassword(RabbitPass);
        final Connection connection = factory.newConnection();
        final Channel channel = connection.createChannel();

        channel.queueDeclare(TASK_QUE_NAME, true, false, false, null);
        System.out.println("[*] Waiting for messages");

        final Consumer consumer = new DefaultConsumer(channel) {
            @Override
            public void handleDelivery(String consumerTag, Envelope envelope, AMQP.BasicProperties properties, byte[] body) throws IOException {
                String message = new String(body, "UTF-8");

                System.out.println(" [x] Received '" + message + "'");
                try {
                    InsertError(message, properties.getHeaders());
                } finally {
                    System.out.println(" [x] Done");
                    channel.basicAck(envelope.getDeliveryTag(), false);
                }
            }

        };
        boolean autoAck = false;
        channel.basicConsume(TASK_QUE_NAME, autoAck, consumer);

    }

    private static void InsertError(String message, Map<String, Object> headers) {

        try {

            String myDriver = "org.gjt.mm.mysql.Driver";
            String myUrl = "jdbc:mysql://"+DBIP+"/test";

            Class.forName(myDriver);
            java.sql.Connection conn = DriverManager.getConnection(myUrl, DBUser, DBPass);

            String query = "INSERT INTO HackerNewsDB.Errors (Post, Location, Error) VALUES (?, ?, ?)";


            PreparedStatement stmt = conn.prepareStatement(query);
            stmt.setString(1, message);
            stmt.setString(2, String.valueOf(headers.get("Error-in")));
            stmt.setString(3, String.valueOf(headers.get("Error-Message")));

            stmt.execute();

            conn.close();

        }catch (Exception e){
            System.err.println("Got an exception!: "+e.getMessage());
            System.exit(0);
        }

    }


}
