---
id: 01KBDHAN6369YHYZTZGB8KKN6T
slug: "edge-computing--en-wikipedia-org-1"
title: "Source: Edge computing - Wikipedia"
url: "https://en.wikipedia.org/wiki/Edge_computing"
type: source
related_article: "edge-computing"
created: 2025-12-02T00:00:00Z
tags: ["topic:edge-computing", "topic:fog-computing", "topic:cloud-computing", "topic:internet-of-things-(iot)", "topic:smart-cities", "topic:serverless-architecture", "topic:content-delivery-network", "topic:mobile-edge-computing", "topic:distributed-computing-architecture", "topic:load-balancing-(computing)", "topic:autonomous-driving", "topic:facial-recognition", "topic:augmented-reality", "topic:industry-4.0", "topic:home-automation", "topic:satellite-systems", "topic:edge-ai", "topic:edge-intelligence", "topic:on-device-ai", "topic:heterogeneous-computing", "topic:wireless-sensor-networks", "topic:radio-frequency-identification", "topic:internet-refrigerator", "topic:edge-computing"]
people: ["person:karim-arabi", "person:alex-reznik"]
summary: "Summarized source material for Edge Computing"
model: "code-prefilter"
language: "en"
---

# Edge Computing

- Please don't skip this 1-minute read. It's Tuesday, 2 December, and if you're like us, you've used Wikipedia countless times. To settle an argument with a friend. To satisfy a curiosity. Whether it's 3 in the morning or afternoon, Wikipedia is useful in your life. Please give $2.75.

- Wikipedia's been around since 2001. Back then, it was just a wildly ambitious, probably impossible dream. But it came together piece by piece—created by people, not machines. Wikipedia's not perfect, but it's always been free thanks to everyday readers.

- Only 2% ever donate. But that small group makes a big difference. When you support Wikipedia, you're standing up for something simple but profound: that knowledge should belong to everyone.

- If you agree, then this is your moment to give back. Even $2.75 makes a difference. Help keep it going—for you, for the next reader, and for the next generation.

- 25 years of the internet at its best

- Thank you for your support!

- Unlock tax benefits by directing your donation via your Donor-Advised Fund (DAF)

- Qualified Charitable Distributions from a tax efficient eligible IRA

- Involve your employer and increase the impact of your donation

- Nearly half of our budget goes toward supporting the technology that powers Wikipedia and other Wikimedia projects.
- We are constantly working to enhance the user experience for both contributors and readers, improve site security, and ensure reliable access to our websites globally.
- This infrastructure and product support sustain one of the top ten most visited websites in the world, all at a fraction of the cost of popular for-profit websites.

- The global reach of Wikimedia projects is made possible by the hard work of volunteers from across the globe.
- We provide grants, legal support, and other resources to help build vibrant volunteer communities.
- Additionally, we promote community engagement through outreach initiatives and advocate for the growth and protection of free knowledge.

- Operational costs are essential for the smooth management and governance of the Wikimedia Foundation. These expenses help us recruit top talent and support staff around the world, empowering them to carry out the mission of the Wikimedia Foundation.

- Donor support is crucial to sustaining Wikipedia and our other free knowledge endeavors. Our team is committed to efficient and effective fundraising throughout the year, ensuring that every contribution helps advance our mission.

- 2.1 Privacy and security

- Edge computing is a distributed computing model that brings computation and data storage closer to the sources of data.
- More broadly, it refers to any design that pushes computation physically closer to a user, so as to reduce the latency compared to when an application runs on a centralized data center.[1]

- The term began being used in the 1990s to describe content delivery networks—these were used to deliver website and video content from servers located near users.[2] In the early 2000s, these systems expanded their scope to hosting other applications,[3] leading to early edge computing services.[4] These services could do things like find dealers, manage shopping carts, gather real-time data, and place ads.

- The Internet of Things (IoT), where devices are connected to the internet, is often linked with edge computing.[5]

- Edge computing involves running computer programs that deliver quick responses close to where requests are made.
- Karim Arabi, during an IEEE DAC 2014 keynote[6] and later at an MIT MTL Seminar in 2015, described edge computing as computing that occurs outside the cloud, at the network's edge, particularly for applications needing immediate data processing.[7]

- Edge computing is often equated with fog computing, particularly in smaller setups.[8] However, in larger deployments, such as smart cities, fog computing serves as a distinct layer between edge computing and cloud computing, with each layer having its own responsibilities.[9][10]

- "The State of the Edge" report explains that edge computing focuses on servers located close to the end-users.[11] Alex Reznik, Chair of the ETSI MEC ISG standards committee, defines 'edge' loosely as anything that's not a traditional data center.[12]

- In cloud gaming, edge nodes, known as "gamelets", are typically within one or two network hops from the client, ensuring quick response times for real-time games.[13]

- Edge computing might use virtualization technology to simplify deploying and managing various applications on edge servers.[14]

- In 2018, the world's data was expected to grow 61 percent to 175 zettabytes by 2025.[15] According to research firm Gartner, around 10 percent of enterprise-generated data is created and processed outside a traditional centralized data center or cloud.
- By 2025, the firm predicts that this figure will reach 75 percent.[16] The increase in IoT devices at the edge of the network is producing a massive amount of data — storing and using all that data in cloud data centers pushes network bandwidth requirements to the limit.[17] Despite the improvements in network technology, data centers cannot guarantee acceptable transfer rates and response times, which often is a critical requirement for many applications.[18] Furthermore, devices at the edge constantly consume data coming from the cloud, forcing companies to decentralize data storage and service provisioning, leveraging physical proximity to the end user.

- In a similar way, the aim of edge computing is to move the computation away from data centers towards the edge of the network, exploiting smart objects, mobile phones, or network gateways to perform tasks and provide services on behalf of the cloud.[19] By moving services to the edge, it is possible to provide content caching, service delivery, persistent data storage, and IoT management resulting in better response times and transfer rates.
- At the same time, distributing the logic to different network nodes introduces new issues and challenges.[20]

- The distributed nature of this paradigm introduces a shift in security schemes used in cloud computing.
- In edge computing, data may travel between different distributed nodes connected via the internet, and thus requires special encryption mechanisms independent of the cloud.
- This approach minimizes latency, reduces bandwidth consumption, and enhances real-time responsiveness for applications.
- Edge nodes may also be resource-constrained devices, limiting the choice in terms of security methods.
- Moreover, a shift from centralized top-down infrastructure to a decentralized trust model is required.[21] On the other hand, by keeping and processing data at the edge, it is possible to increase privacy by minimizing the transmission of sensitive information to the cloud.
- Furthermore, the ownership of collected data shifts from service providers to end-users.[22]

- Scalability in a distributed network must face different issues.
- First, it must take into account the heterogeneity of the devices, having different performance and energy constraints, the highly dynamic condition, and the reliability of the connections compared to more robust infrastructure of cloud data centers.
- Moreover, security requirements may introduce further latency in the communication between nodes, which may slow down the scaling process.[18]

- The state-of-the-art scheduling technique can increase the effective utilization of edge resources and scales the edge server by assigning minimum edge resources to each offloaded task.[23]

- Management of failovers is crucial in order to keep a service alive.
- If a single node goes down and is unreachable, users should still be able to access a service without interruptions.
- Moreover, edge computing systems must provide actions to recover from a failure and alert the user about the incident.
- To this aim, each device must maintain the network topology of the entire distributed system, so that detection of errors and recovery become easily applicable.
- Other factors that may influence this aspect are the connection technologies in use, which may provide different levels of reliability, and the accuracy of the data produced at the edge that could be unreliable due to particular environment conditions.[18] As an example, an edge computing device, such as a voice assistant, may continue to provide service to local users even during cloud service or internet outages.[22]

- Edge computing brings analytical computational resources close to the end users and therefore can increase the responsiveness and throughput of applications.
- A well-designed edge platform would significantly outperform a traditional cloud-based system.
- Some applications rely on short response times, making edge computing a significantly more feasible option than cloud computing.
- Examples range from IoT to autonomous driving,[24] anything health or human / public safety relevant,[25] or involving human perception such as facial recognition, which typically takes a human between 370-620 ms to perform.[26] Edge computing is more likely to be able to mimic the same perception speed as humans, which is useful in applications such as augmented reality, where the headset should preferably recognize who a person is at the same time as the wearer does.

- Due to the nearness of the analytical resources to the end users, sophisticated analytical tools and artificial intelligence tools can run on the edge of the system. This placement at the edge helps to increase operational efficiency and is responsible for many advantages to the system.

- Additionally, the usage of edge computing as an intermediate stage between client devices and the wider internet results in efficiency savings that can be demonstrated in the following example: A client device requires computationally intensive processing on video files to be performed on external servers.
- By using servers located on a local edge network to perform those computations, the video files only need to be transmitted in the local network.
- Avoiding transmission over the internet results in significant bandwidth savings and therefore increases efficiency.[26] Another example is voice recognition.
- If the recognition is performed locally, it is possible to send the recognized text to the cloud rather than audio recordings, significantly reducing the amount of required bandwidth.[22]

- Edge application services reduce the volumes of data that must be moved, the consequent traffic, and the distance that data must travel.
- That provides lower latency and reduces transmission costs.
- Computation offloading for real-time applications, such as facial recognition algorithms, showed considerable improvements in response times, as demonstrated in early research.[27] Further research showed that using resource-rich machines called cloudlets or micro data centers near mobile users, which offer services typically found in the cloud, provided improvements in execution time when some of the tasks are offloaded to the edge node.[28] On the other hand, offloading every task may result in a slowdown due to transfer times between device and nodes, so depending on the workload, an optimal configuration can be defined.

- An IoT-based power grid system enables communication of electricity and data to monitor and control the power grid,[29] which makes energy management more efficient.

- Other notable applications include connected cars, self-driving cars,[30] smart cities,[31] Industry 4.0, home automation,[32] missiles,[33] and satellite systems.[34] The growing field of edge artificial intelligence (edge AI or edge intelligence, sometimes referred to as "local AI" or "on-device AI") implements artificial intelligence in an edge computing environment, on the device or close to where data is collected.[35][36]

- Content delivery network

- Edge data integration

- Heterogeneous computing

- Mobile edge computing

- Serverless architecture

- Board support package

- Single-board computer Raspberry Pi

- Firmware Custom firmware Proprietary firmware

- Hacking of consumer electronics

- Homebrew (video games)

- PlayStation 3 Jailbreak

- Linux on embedded systems

- Linux for mobile devices

- Light-weight Linux distribution

- Real-time operating system

- List of open-source hardware

- Wireless sensor networks

- Radio-frequency identification

- Internet refrigerator

- Josef Preishuber-Pflügl

- Post-cloud computing architecture

- Distributed computing architecture

- Load balancing (computing)

- All articles with dead external links

- Articles with dead external links from February 2025

- Articles with dead external links from May 2025

- Articles with permanently dead external links

- CS1 Serbian-language sources (sr)

- CS1 errors: missing periodical

- Articles with short description

- Short description is different from Wikidata

- Edit preview settings
