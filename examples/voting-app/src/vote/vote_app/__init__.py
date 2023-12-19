import json
import random

from flask import Flask, render_template, request, make_response, g
from redis import Redis

from config import Config


# retrieve redis connection from g object if not exists
# read more about g object here: https://flask.palletsprojects.com/en/2.3.x/appcontext/
def get_redis():
    if 'redis' not in g:
        g.redis = Redis(host=Config.REDIS_HOST, port=Config.REDIS_PORT, socket_timeout=5)
    return g.redis


# create app application factory. more details here: https://flask.palletsprojects.com/en/2.3.x/patterns/appfactories/
def create_app(config=Config):

    app = Flask(__name__)
    app.config.from_object(config)

    @app.route("/", methods=['POST', 'GET'])
    def index():
        voter_id = request.cookies.get('voter_id')
        if not voter_id:
            voter_id = hex(random.getrandbits(64))[2:-1]

        vote = None

        if request.method == 'POST':
            vote = request.form['vote']
            app.logger.info('Received vote for %s', vote)
            data = json.dumps({'voter_id': voter_id, 'vote': vote})
            get_redis().rpush('votes', data)

        resp = make_response(render_template(
            'index.html',
            option_a=config.OPTION_A,
            option_b=config.OPTION_B,
            hostname=config.HOSTNAME,
            vote=vote,
        ))

        resp.set_cookie('voter_id', voter_id)

        return resp

    return app
