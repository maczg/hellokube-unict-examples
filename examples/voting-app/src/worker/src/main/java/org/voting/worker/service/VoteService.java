package org.voting.worker.service;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.voting.worker.model.Vote;
import org.voting.worker.repository.VoteRepository;

@Service
public class VoteService {

    @Autowired
    private VoteRepository voteRepository;

    public Vote getVote(String voterId) {//vedo se voterId ha già votato e ritorno il vote altrimenti null
        Vote avote= voteRepository.findByVoterId(voterId);
        if (avote == null) return null;
        return avote;
    } //Long voterId

    public void saveVote(String voterId, String vote) {
        if ((voteRepository.findByVoterId(voterId))!=null) {
            System.out.println("save new vote value " +  vote + " of voter_id= " + voterId);
            deleteVote(voterId);
            //voteRepository.findByVoterId(voterId).setVote(vote);
        }
        voteRepository.save(new Vote(voterId,vote));
    }


    public void deleteVote(String voterId) {//Long voterId
        voteRepository.delete(voteRepository.findByVoterId(voterId));
    }


}
