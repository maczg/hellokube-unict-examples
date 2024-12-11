package org.voting.worker.service;


import org.json.JSONObject;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.scheduling.annotation.Scheduled;
import org.springframework.stereotype.Component;

import java.util.List;

@Component
public class WorkerPeriodicTask {

    @Value("${redis.addr}")
    private String redisAddress;

    @Autowired
    VoteService voteService;

    @Scheduled(fixedRate = 5000)
    public void processVote() {
        RedisService redisService = new RedisService(redisAddress, 6379);
        redisService.connect();
        List<String> voteData = (List<String>) redisService.getVote();
        System.out.println("voteData= " + voteData);
        String voteJsonString= voteData.get(1);
        JSONObject voteJson = new JSONObject(voteJsonString);
        System.out.println("voteJsonString= " + voteJsonString);
        System.out.println("voteJson" + voteJson);
        String voterId = voteJson.getString("voter_id");
        String vote = voteJson.getString("vote");
        System.out.println("voterId = " + voterId + ", vote = " + vote);
        voteService.saveVote(voterId,vote);
    }
}
