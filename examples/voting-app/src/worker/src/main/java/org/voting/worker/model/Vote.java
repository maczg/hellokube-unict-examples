package org.voting.worker.model;

import jakarta.persistence.*;

@Entity
public class Vote {

    @Id
    @GeneratedValue(strategy=GenerationType.AUTO)
    private Long id;

    @Column(unique = true)
    private String voterId;

    private String vote;

    public Vote() {}

    public Vote(String voterId, String vote) {
        this.voterId = voterId;
        this.vote = vote;
    }

    public Long getId() {
        return id;
    }

    public String getVoterId() {
        return voterId;
    }

    public String getVote() {
        return vote;
    }

    public Vote setVoterId(String voterId) {
        this.voterId = voterId;
        return this;
    }

    public Vote setVote(String vote) {
        this.vote = vote;
        return this;
    }

    @Override
    public String toString() {
        return "Vote{" +
                "id=" + id +
                ", voterId=" + voterId +
                ", vote='" + vote + '\'' +
                '}';
    }
}
