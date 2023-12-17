import os

from flask import Flask, render_template, request, make_response, g
from pymongo import MongoClient

app = Flask(__name__)

mongo_user = os.getenv('MONGO_USER', 'admin')
mongo_password = os.getenv('MONGO_PASSWORD', 'password')
mongo_host = os.getenv('MONGO_HOST', 'localhost')
mongo_port = os.getenv('MONGO_PORT', '27017')
mongo_db = os.getenv('MONGO_DB', 'votes')


client = MongoClient(mongo_host, int(mongo_port), username=mongo_user, password=mongo_password)

db = client.votes
votes_collection = db.votes

pipeline = [
    {
        '$group': {
            '_id': '$vote',  # Group by the 'vote' field
            'count': {'$sum': 1}  # Count the number of documents in each group
        }
    }
]


@app.route("/results", methods=['POST', 'GET'])
def hello():
    results = list(votes_collection.aggregate(pipeline))
    return results


if __name__ == '__main__':
    app.run(host='0.0.0.0', debug=True, port=5001)
