FROM redis:alpine

# Expose the Redis port (default is 6379)
EXPOSE 6379

COPY redis.conf /usr/local/etc/redis/redis.conf
# Run Redis server
CMD ["redis-server", "/usr/local/etc/redis/redis.conf"]