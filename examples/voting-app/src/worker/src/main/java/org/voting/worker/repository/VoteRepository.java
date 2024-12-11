package org.voting.worker.repository;


import org.voting.worker.model.Vote;
import org.springframework.data.repository.CrudRepository;

public interface VoteRepository extends CrudRepository<Vote, Long> {
     Vote findByVoterId(String voterId);
}
