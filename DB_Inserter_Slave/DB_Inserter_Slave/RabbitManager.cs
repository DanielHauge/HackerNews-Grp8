using RabbitMQ.Client;
using RabbitMQ.Client.Events;
using System;
using System.Text;
using System.Web.Script.Serialization;
using MySql.Data.MySqlClient;

namespace DB_Inserter_Slave
{
    class RabbitManager
    {
        private MySqlConnection sqlConnection { get; set; }

        public RabbitManager()
        {
            sqlConnection = new MySqlConnection("server = 46.101.103.163; user id = admin; Password=password; database = HackerNewsDB; allowuservariables = True; persistsecurityinfo = True");
        }
        public void InsertMessage(string message)
        {
            MySqlCommand command = new MySqlCommand(message, sqlConnection);
            MySqlDataReader reader;
            sqlConnection.Open();
            reader = command.ExecuteReader();
            while (reader.Read())
            {

            }
            sqlConnection.Close();
        }

        public void ReceiveMessage(string messageChannel)
        {
            var factory = new ConnectionFactory() { HostName = "138.197.186.82", UserName = "admin", Password = "password" };
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
                    //decide where the message go
                    if (messageChannel == "ThreadInsert")
                    {
                        var threadMessage = (JsonMessage)new JavaScriptSerializer().Deserialize(Encoding.UTF8.GetString(body), typeof(JsonMessage));
                        MySqlCommand command = new MySqlCommand("Select ID from HackerNewsDB.User where Name = '" + threadMessage.username + "'", sqlConnection);
                        MySqlDataReader reader;
                        sqlConnection.Open();
                        reader = command.ExecuteReader();
                        int userID = 0;
                        while (reader.Read())
                        {
                            string result = reader[0].ToString();
                            userID = int.Parse(result);
                        }
                        sqlConnection.Close();
                        Threads thread = new Threads { Name = threadMessage.post_title, UserID = userID, Post_URL = threadMessage.post_url, Han_ID = threadMessage.harnesst_id, Time = DateTime.Now };
                        string message = "Insert into HackerNewsDB.Thread(Name,UserID,Time,Han_ID,PostURL) values('" + thread.Name + "','" + thread.UserID + "','" + thread.Time + "','" + thread.Han_ID + "','" + thread.Post_URL + "')";
                        Console.WriteLine("Thread get");
                        InsertMessage(message);
                    }
                    else if (messageChannel == "HNPost")
                    {
                        var commentMessage = (JsonMessage)new JavaScriptSerializer().Deserialize(Encoding.UTF8.GetString(body), typeof(JsonMessage));
                        MySqlCommand command = new MySqlCommand("Select ID from HackerNewsDB.User where Name = '" + commentMessage.username + "'", sqlConnection);
                        MySqlDataReader reader;
                        sqlConnection.Open();
                        reader = command.ExecuteReader();
                        int userID = 0;
                        while (reader.Read())
                        {
                            string result = reader[0].ToString();
                            userID = int.Parse(result);
                        }
                        sqlConnection.Close();
                        Comment comment = new Comment { UserID = userID, ThreadID = commentMessage.post_parent, ParentID = commentMessage.post_parent, Han_ID = commentMessage.harnesst_id, Name = commentMessage.post_text, Time = DateTime.Now };
                        string message = "Insert into HackerNewsDB.Comment(ThreadID,Name,UserID,Time) values('" + comment.ParentID + "','" + comment.Name + "','" + comment.UserID + "','" + comment.Time + "')";
                        Console.WriteLine("Comment get");
                        InsertMessage(message);
                    }
                    else if (messageChannel == "UserInsert")
                    {
                        //var userMessage = (JsonMessage)new JavaScriptSerializer().Deserialize(Encoding.UTF8.GetString(body), typeof(JsonMessage));
                        //string message = "Insert into HackerNewsDB.User(Name,KarmaPoints,Password,Email) values('" + userMessage.Name + "','" + userMessage.KarmaPoints + "','" + userMessage.Password + "','" + userMessage.Email + "')";
                        //Console.WriteLine("User get");
                        //InsertMessage(message);
                    }
                    else
                    {
                        Console.WriteLine("ERROR");
                    }

                    Console.WriteLine(" [x] Done");

                    channel.BasicAck(deliveryTag: ea.DeliveryTag, multiple: false);
                };
                channel.BasicConsume(queue: messageChannel,
                                     autoAck: false,
                                     consumer: consumer);

            }
        }
    }
}
