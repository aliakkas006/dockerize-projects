from flask import Flask

def create_app():
    app = Flask(__name__)

    @app.route('/')
    def home():
        return {"message": "Hello from Flask API using Docker!"}


    return app
