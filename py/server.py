from concurrent.futures import ThreadPoolExecutor
from summarization_pb2 import SummaryResponse
from summarization_pb2_grpc import SummarizationServicer, add_SummarizationServicer_to_server
import logging
import grpc 
import torch
import warnings
warnings.filterwarnings(action = 'ignore')
from transformers import AutoTokenizer, AutoModelWithLMHead

class ModelInit:
    def __init__(self, model_name = 't5-base'):
        self.tokenizer = AutoTokenizer.from_pretrained('t5-base')
        self.model = AutoModelWithLMHead.from_pretrained('t5-base')

class SummarizationService(SummarizationServicer):
    def __init__(self):
        pass
    def GetSummary(self, request):
        logging.info(f"Full Text" + self.request)
        text = self.request
        t5Model = ModelInit()
        tokenizer = t5Model.tokenizer
        model = t5Model.model
        inputs = tokenizer.encode(
            "summarize: " + text,
            return_tensors = 'pt',
            max_length = 512,
            truncation = True 
        )
        summary_ids = model.generate(
            inputs, 
            max_length = 150,
            min_length = 80,
            length_penalty=5., num_beams=2
        )

        summary = tokenizer.batch_decode(summary_ids[0], skip_special_tokens=True, clean_up_tokenization_spaces=True)

        resp = SummaryResponse(response = summary)

        return resp



if __name__ == "__main__":
    logging.basicConfig(
        level = logging.INFO,
        format='%(asctime)s - %(levelname)s - %(message)s',
    )
    server = grpc.server(ThreadPoolExecutor())
    add_SummarizationServicer_to_server(SummarizationService, server)
    port = 8080
    server.add_insecure_port(f'[::]:{port}')
    server.start()
    logging.info(f'Server Ready on Port {port}')
    server.wait_for_termination()



        
