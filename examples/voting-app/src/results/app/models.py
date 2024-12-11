from app.database import db

class Vote(db.Model):
    __tablename__ = 'vote'

    id = db.Column(db.BigInteger, primary_key=True)
    voter_id = db.Column(db.String(100), unique=True, nullable=False)
    vote = db.Column(db.String(10), nullable=False)