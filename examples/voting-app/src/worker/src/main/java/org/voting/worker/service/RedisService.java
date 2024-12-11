package org.voting.worker.service;

import redis.clients.jedis.Jedis;
import java.util.logging.Level;
import java.util.logging.Logger;

import static java.lang.Thread.sleep;

public class RedisService {
    private Jedis redis;
    private static final Logger logger = Logger.getLogger(RedisService.class.getName());


    public RedisService(String hostname, int port) {
        this.redis = new Jedis(hostname, port);
    }

    public void connect() {
        int maxRetries = 10;
        int retries = 0;
        while (retries < maxRetries){
            try {
                logger.log(Level.INFO, "Attempting to connect to Redis, try {0}", retries + 1);
                this.redis.connect();
                logger.log(Level.INFO, "Successfully connected to Redis");
                break;

            } catch (Exception e){
                logger.log(Level.WARNING, "Failed to connect to Redis, retrying...");
                retries++;
                try {
                    sleep(10000);
                } catch (InterruptedException ex) {
                    logger.log(Level.SEVERE, "Thread interrupted", ex);
                }
            }
            if (retries == maxRetries) {
                logger.log(Level.SEVERE, "Exceeded maximum retries, could not connect to Redis");
            }
        }
    }

    public Iterable<String> getVote(){
        return this.redis.blpop(0, "votes");
    }

    public String getVoteString(){
        return this.redis.blpop(0, "votes").toString();
    }
}