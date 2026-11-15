# Roadmap

## Vision 
A meta search aggregator that feed the indiced web to llm inference or harness ...  for further navigation and retrival from web urls.

### Goal :
- To provide slimmest possible binary with low runtime resource usage that runs out of the box without any dependency.
- Ideal for resource constraint systems , Edge AI and local harnesses.
- A socket based connection interface ... for lowest overhead for ipc and fastest possible data transfer ... thus is embeddable and integrated in any ecosystem.

## Achievements
- We forked a repo [SearxGo](https://github.com/Openmentia/SearxGo) , as the project is small and easily navigateable. Building on top is better than , from scratch.
- Stripped http server and result serving GUI interface.
- Now , a cli interface outputs json results ... a good data shape for AI harnesses


# Future Goals : to do's
- A socket based service/interface : what it does it takes query over socket connection and returns the json 
- New engines : brave search , or other content heavy / search engines
- Optimisation : removing unecessary code and making the codebase simple
- Further will be added ... on the way
