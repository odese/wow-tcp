# WoW - TCP
A TCP server and client example for word of wisdom, with Proof of Work ( challenge-response protocol ) integration in order to protect server from DDOS attacks.
## Chosen PoW Algorithm:
Despite the fact that Hashcash is the most popular and preferable pow-algorithm in the industry, I chose a different algorithm which is finding the two big prime divisors of another big integer. There are not any particular reason for this choice. Probably using hashcash would be better, but after all, they are all algorithms and they all have some pros and cons. 

One credible reason to choose this algorithm might be the ability to harden or ease the challange by your preferences (e.g Adding more prime numbers or using much more bigger numbers to create a challange or vice versa.) since you have direct control on creating a challange. In hashcash, creator doesn't have that impact on the challange.

Another reason for me is that, well, I studied mathematics and cryptography in university and I haven't use any numerical manipulations on codes for years, since then. And wanted to try something different from the what the software industry's choices.

As I mentioned before, all those algoritms are one-to-one mathematical problems for computers to solve, while one way is hard and other way is easy. Multiplication is easier than division.

## Protocol Flow
![550px-Proof_of_Work_challenge_response svg](https://user-images.githubusercontent.com/46742766/177880500-9d087d2d-e41f-46e5-85d0-9bc2a4440d8d.png)

## Project Structure
![project structure](https://user-images.githubusercontent.com/46742766/177880198-cb7cbc7f-11e3-4ff1-bd80-dd88cf04f807.png)

Project is tried to be structured and developed as Standard Go Project Layout with Clean Code Architecture.

## Example Output
![output](https://user-images.githubusercontent.com/46742766/177882440-b14476b7-2a9d-4d08-9d3e-20d40c77c937.png)

## How to Run
### Server:
```
cd server
```
```
go run cmd\main.go
```
### Client:
```
cd client
```
```
go run cmd\main.go
```

Also both server and client have a dockerfile in their build folders.
