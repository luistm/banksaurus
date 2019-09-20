from invoke import task


@task
def run_containers(c):
    c.run("docker-compose -f deployment/docker-compose.json up -d")


@task
def stop_containers(c):
    c.run("docker-compose -f deployment/docker-compose.json down")
