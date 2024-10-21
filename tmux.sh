#!/bin/sh

TMUX_SESSION="lizardpoint"

tmux has-session -t $TMUX_SESSION 2> /dev/null

if [ $? != 0 ]; then
	tmux new-session -d -s $TMUX_SESSION -n "term"
	tmux send-keys -t $TMUX_SESSION:term "cd ~/Projects/lizardpoint" C-m
	tmux send-keys -t $TMUX_SESSION:term "codium ." C-m
	tmux send-keys -t $TMUX_SESSION:term "clear" C-m

	tmux new-window -t $TMUX_SESSION -n "docker"
	tmux send-keys -t $TMUX_SESSION:docker "docker-compose up -d --force-recreate" C-m
	tmux send-keys -t $TMUX_SESSION:docker "clear" C-m
	tmux send-keys -t $TMUX_SESSION:docker "docker logs -f lizardpoint-web" C-m

	tmux select-window -t $TMUX_SESSION:term
fi

tmux attach-session -t $TMUX_SESSION
