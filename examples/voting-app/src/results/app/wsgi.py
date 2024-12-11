import os
from datetime import datetime

from flask import Flask, jsonify
from app.database import db
from app.models import Vote
from sqlalchemy.sql import func
from sqlalchemy.exc import SQLAlchemyError

MYSQL_USERNAME = os.environ['MYSQL_USER']
MYSQL_PASSWORD = os.environ['MYSQL_PASSWORD']
MYSQL_HOST = os.environ['MYSQL_HOST']
MYSQL_PORT = os.environ['MYSQL_PORT']
MYSQL_DATABASE = os.environ['MYSQL_DB']

app = Flask(__name__)
app.config['SQLALCHEMY_DATABASE_URI'] = f'mysql+pymysql://{MYSQL_USERNAME}:{MYSQL_PASSWORD}@{MYSQL_HOST}:{MYSQL_PORT}/{MYSQL_DATABASE}'

db.init_app(app)
@app.route("/api/percentages", methods=["GET"])
def calculate_vote_percentages():
    try:
        query = db.session.query(Vote.vote, func.count(Vote.vote).label("count")).group_by(Vote.vote)
        results = query.all()

        total_votes = sum(row.count for row in results)
        percentages = {
            "cat": round(next((row.count for row in results if row.vote == "a"), 0) / total_votes * 100, 2) if total_votes > 0 else 0,
            "dog": round(next((row.count for row in results if row.vote == "b"), 0) / total_votes * 100, 2) if total_votes > 0 else 0,
        }

        timestamp = datetime.now().strftime("%Y-%m-%d %H:%M:%S")

        return jsonify({
            "timestamp": timestamp,
            "percentages": percentages,
        })

    except SQLAlchemyError as e:
        return jsonify({"error": str(e)}), 500

if __name__ == '__main__':
    app.run(host='0.0.0.0', debug=True)
