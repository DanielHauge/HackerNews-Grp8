using RabbitMQ.Client;
using RabbitMQ.Client.Events;
using System;
using System.Text;
using System.Web.Script.Serialization;
using MySql.Data.MySqlClient;
using System.Collections.Generic;

namespace DB_Inserter_Slave
{
    class RabbitManager
    {
        private string sqlString { get; set; }

        public RabbitManager()
        {
            sqlString = "server = " + Program.dbip + "; user id = " + Program.dbusername + "; password = " + Program.dbpassword + "; database = HackerNewsDB; allowuservariables = True; persistsecurityinfo = True";
        }
        public void InsertMessage(MySqlCommand command)
        {
            MySqlConnection sqlConnection = new MySqlConnection(sqlString);
            sqlConnection.Open();
            command.Connection = sqlConnection;
            command.ExecuteNonQuery();
            sqlConnection.Close();
            sqlConnection.Dispose();
        }
        
        public void ReceiveMessage(string messageChannel)
        {
            var factory = new ConnectionFactory() { HostName = Program.rabbitip, UserName = Program.rabbituser, Password = Program.rabbitpassword };
            using (var connection = factory.CreateConnection())
            using (var channel = connection.CreateModel())
            {
                channel.QueueDeclare(queue: messageChannel,
                                     durable: true,
                                     exclusive: false,
                                     autoDelete: false,
                                     arguments: null);

                channel.BasicQos(0, 1, false);

                Console.WriteLine(" [*] Waiting for messages.");

                var consumer = new EventingBasicConsumer(channel);
                consumer.Received += (model, ea) =>
                {
                    var body = ea.Body;
                    try
                    {
                        var jsonmessage = (JsonMessage)new JavaScriptSerializer().Deserialize(Encoding.UTF8.GetString(body), typeof(JsonMessage));
                        //decide where the message go

                        if (jsonmessage.post_type == "story")
                        {
                            MySqlConnection sqlConnection = new MySqlConnection(sqlString);
                            MySqlCommand command = new MySqlCommand("Select ID from HackerNewsDB.User where Name LIKE '" + jsonmessage.username + "';", sqlConnection);
                            sqlConnection.Open();
                            MySqlDataReader reader = command.ExecuteReader();
                            int userID = 0;
                            while (reader.Read())
                            {
                                string result = reader[0].ToString();
                                userID = int.Parse(result);
                            }
                            sqlConnection.Close();
                            sqlConnection.Dispose();
                            Threads thread = new Threads { Name = jsonmessage.post_title, UserID = userID, Post_URL = jsonmessage.post_url, Han_ID = jsonmessage.hanesst_id, Time = DateTime.Now };
                            //string message = "Insert into HackerNewsDB.Thread(Name,UserID,Time,Han_ID,Post_URL) values(" + thread.Name + "," + thread.UserID + "," + thread.Time + "," + thread.Han_ID + "," + thread.Post_URL + ")";
                            MySqlCommand InsertCommand = new MySqlCommand("Insert into HackerNewsDB.Thread (Name, UserID, Time, Han_ID, Post_URL, Karma) values (@Name, @UserID, @Time, @Han_ID, @Post_URL, @Karma)");
                            InsertCommand.Parameters.AddWithValue("@Name", thread.Name);
                            InsertCommand.Parameters.AddWithValue("@UserID", thread.UserID);
                            InsertCommand.Parameters.AddWithValue("@Karma", 0);
                            InsertCommand.Parameters.AddWithValue("@Time", thread.Time);
                            InsertCommand.Parameters.AddWithValue("@Han_ID", thread.Han_ID);
                            InsertCommand.Parameters.AddWithValue("@Post_URL", thread.Post_URL);
                            Console.WriteLine("Thread get");
                            InsertMessage(InsertCommand);
                        }
                        else if (jsonmessage.post_type == "comment")
                        {
                            MySqlConnection sqlConnection = new MySqlConnection(sqlString);
                            MySqlCommand command = new MySqlCommand("Select ID from HackerNewsDB.User where Name LIKE '" + jsonmessage.username + "';", sqlConnection);
                            sqlConnection.Open();
                            MySqlDataReader reader = command.ExecuteReader();
                            int userID = 0;
                            while (reader.Read())
                            {
                                string result = reader[0].ToString();
                                userID = int.Parse(result);
                            }
                            sqlConnection.Close();
                            sqlConnection.Dispose();
                            Comment comment;
                            if (jsonmessage.hanesst_id > 0)
                            {
                                Console.WriteLine("Woops the han_id was above 0 = simulator insert!");
                                MySqlConnection sqlConnection1 = new MySqlConnection(sqlString);
                                MySqlCommand command1 = new MySqlCommand("Select ID from HackerNewsDB.Thread where Han_ID LIKE '" + jsonmessage.post_parent + "';", sqlConnection1);
                                sqlConnection1.Open();
                                MySqlDataReader reader1 = command1.ExecuteReader();
                                int RealThreadID = 0;
                                while (reader1.Read())
                                {
                                    string result = reader1[0].ToString();
                                    RealThreadID = int.Parse(result);

                                }
                                sqlConnection1.Close();
                                sqlConnection1.Dispose();

                                if (RealThreadID == 0)
                                {
                                    Console.WriteLine("Woops! the ThreadID was not found = comment is a comment of another comment");
                                    MySqlConnection sqlConnection2 = new MySqlConnection(sqlString);
                                    MySqlCommand command2 = new MySqlCommand("Select ThreadID from HackerNewsDB.Comment where Han_ID LIKE '" + jsonmessage.post_parent + "';", sqlConnection2);
                                    sqlConnection2.Open();
                                    MySqlDataReader reader2 = command2.ExecuteReader();
                                    while (reader2.Read())
                                    {
                                        string result = reader2[0].ToString();
                                        RealThreadID = int.Parse(result);
                                    }
                                    sqlConnection2.Close();
                                    sqlConnection2.Dispose();


                                }
                                sqlConnection1.Close();
                                sqlConnection1.Dispose();


                                comment = new Comment { UserID = userID, ThreadID = RealThreadID, Name = jsonmessage.post_text, Han_ID = jsonmessage.hanesst_id, Time = DateTime.Now, ParentID = jsonmessage.post_parent };
                            }
                            else
                            {
                                comment = new Comment { UserID = userID, ThreadID = jsonmessage.post_parent, Name = jsonmessage.post_text, Han_ID = jsonmessage.hanesst_id, Time = DateTime.Now };
                            }

                            MySqlCommand InsertCommand = new MySqlCommand("Insert into HackerNewsDB.Comment (ThreadID, Name, UserID, Karma, Time, Han_ID, PostParrent) values (@ThreadID, @Name, @UserID, @Number, @Time, @Han_ID, @PostParrent)");
                            InsertCommand.Parameters.AddWithValue("@Name", comment.Name);
                            InsertCommand.Parameters.AddWithValue("@UserID", comment.UserID);
                            InsertCommand.Parameters.AddWithValue("@Number", 0);
                            InsertCommand.Parameters.AddWithValue("@Time", comment.Time);
                            InsertCommand.Parameters.AddWithValue("@Han_ID", comment.Han_ID);
                            InsertCommand.Parameters.AddWithValue("@ThreadID", comment.ThreadID);
                            InsertCommand.Parameters.AddWithValue("@PostParrent", comment.ParentID);
                            Console.WriteLine("Comment get");
                            InsertMessage(InsertCommand);
                        }
                        else if (jsonmessage.post_type == "UserInsert")
                        {
                            //string message = "Insert into HackerNewsDB.User(Name,KarmaPoints,Password,Email) values('" + userMessage.Name + "','" + userMessage.KarmaPoints + "','" + userMessage.Password + "','" + userMessage.Email + "')";
                            //Console.WriteLine("User get");
                            //InsertMessage(message);
                        }
                        else
                        {
                            Console.WriteLine("ERROR");
                        }
                    }
                    catch (global::System.Exception)
                    {
                        if (ea.BasicProperties.Headers == null)
                        {
                            ea.BasicProperties.Headers = new Dictionary<string, object>();
                            ea.BasicProperties.Headers["loop"] = "no";
                        }
                        else if (ea.BasicProperties.Headers["loop"].ToString() == "yes")
                        {
                            var que = channel.QueueDeclare(queue: "HNError",
                                durable: true,
                                exclusive: false,
                                autoDelete: false,
                                arguments: null);
                            //ea.BasicProperties.ReplyTo = que.QueueName;
                            channel.BasicPublish(exchange: que, routingKey: "", basicProperties: ea.BasicProperties, body: body);
                            //channel.BasicAck(deliveryTag: ea.DeliveryTag, multiple: false);
                            //channel.BasicConsume(queue: messageChannel,
                            //                 autoAck: false,
                            //                 consumer: consumer);
                        }
                        else
                        {
                            var que = channel.QueueDeclare(queue: messageChannel,
                                durable: true,
                                exclusive: false,
                                autoDelete: false,
                                arguments: null);
                            //ea.BasicProperties.ReplyTo = que.QueueName;
                            ea.BasicProperties.Headers["loop"] = "yes";
                            channel.BasicPublish(exchange: que, routingKey: "", basicProperties: ea.BasicProperties, body: body);
                            //channel.BasicAck(deliveryTag: ea.DeliveryTag, multiple: false);
                            //channel.BasicConsume(queue: messageChannel,
                            //                 autoAck: false,
                            //                 consumer: consumer);
                        }
                    }
                    

                    Console.WriteLine(" [x] Done");



                    channel.BasicAck(deliveryTag: ea.DeliveryTag, multiple: false);
                };
                channel.BasicConsume(queue: messageChannel,
                                     autoAck: false,
                                     consumer: consumer);
                Console.ReadLine();
            }
        }

        public void ReceiveUser(string messageChannel)
        {
            var factory = new ConnectionFactory() { HostName = Program.rabbitip, UserName = Program.rabbituser, Password = Program.rabbitpassword };
            using (var connection = factory.CreateConnection())
            using (var channel = connection.CreateModel())
            {
                channel.QueueDeclare(queue: messageChannel,
                                     durable: true,
                                     exclusive: false,
                                     autoDelete: false,
                                     arguments: null);

                channel.BasicQos(0, 1, false);

                Console.WriteLine(" [*] Waiting for User.");

                var consumer = new EventingBasicConsumer(channel);
                consumer.Received += (model, ea) =>
                {
                    var body = ea.Body;
                    var User = (User)new JavaScriptSerializer().Deserialize(Encoding.UTF8.GetString(body), typeof(User));

                    MySqlCommand InsertCommand = new MySqlCommand("Insert into HackerNewsDB.User (Name, Karma, Password, Email) values (@Name, @Karma, @Password, @Email)");
                    InsertCommand.Parameters.AddWithValue("@Name", User.username);
                    InsertCommand.Parameters.AddWithValue("@Karma", 0);
                    InsertCommand.Parameters.AddWithValue("@Password", User.password);
                    InsertCommand.Parameters.AddWithValue("@Email", User.email_addr);
                    Console.WriteLine("User get");
                    InsertMessage(InsertCommand);


                    Console.WriteLine(" [x] Done");



                    channel.BasicAck(deliveryTag: ea.DeliveryTag, multiple: false);
                };
                channel.BasicConsume(queue: messageChannel,
                                     autoAck: false,
                                     consumer: consumer);
                Console.ReadLine();
            }
        }
    }
}
