package main

import (
	"context"
	"log"
	"main/pb"

	"google.golang.org/grpc"
)

func main() {
	addr := "localhost:8080"
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := pb.NewSummarizationClient(conn)
	inp_str := "A theme is any universal idea explored in a literary work. After reading the novel Lolita it became obvious that there were multiple themes occurring throughout the book.In my eyes the most important theme of them all was the power of diction and how Nabokov honored words because they elevated his artwork otherwise dreadful topic. This particular book is known for being risqué, but it is important to note that there are no four-letter words or any obvious graphic material; that's because of Humbert's word choice. The language used in Lolita successfully overcompensates the unadvisable content and allows a sense of beauty to prevail. Subjects such as murder, pedophilia, rape, and even incest are surprisingly appealing due to the way Humbert Humbert narrates each scene with powerful word choice. Humbert uses diction and other forms of diction such as alliteration and imagery to ensure the captivation of his readers, entangling and convincing them into buying his version of the confession."
	req := pb.SummaryRequest{
		Request: inp_str,
	}
	resp, err := client.GetSummary(context.Background(), &req)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Predicted Summary: %v", resp.Response)
}
