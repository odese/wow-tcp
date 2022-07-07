# WoW - TCP
A TCP server and client example for word of wisdom, with Proof of Work ( challenge-response protocol ) integration in order to protect server from DDOS attacks.
## Chosen PoW Algorithm:
Despite the fact that Hashcash is the most popular and preferable pow-algorithm in the industry, I chose a different algorithm which is finding the two big prime divisors of another big integer. There are not any particular reason for this choice. Probably using hashcash would be better, but after all, they are all algorithms and they all have some pros and cons. 

One credible reason to choose this algorithm might be the ability to harden or ease the challange by your preferences (e.g Adding more prime numbers or using much more bigger numbers to create a challange or vice versa.) since you have direct control on creating a challange. In hashcash, creator doesn't have that impact on the challange.

Another reason for me is that, well, I studied mathematics and cryptography in university and I haven't use any numerical manipulations on codes for years, since then. And wanted to try something different from the what the software industry's choices.

As I mentioned before, all those algoritms are one-to-one mathematical problems for computers to solve, while one way is hard and other way is easy. Multiplication is easier than division.

## Protocol Description
<img src="https://upload.wikimedia.org/wikipedia/commons/thumb/5/55/Proof_of_Work_challenge_response.svg/550px-Proof_of_Work_challenge_response.svg.png" width="512"/>
