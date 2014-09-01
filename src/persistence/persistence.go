package persistence

import (
  "time"
  "github.com/fzzy/radix/redis"
)

func getClient() (*redis.Client, error) {
  c, err := redis.DialTimeout("tcp", "127.0.0.1:6379", time.Duration(10)*time.Second)
  if err != nil {
    return c, err
  }

  c.Cmd("select", 0)
  return c, nil
}

func Save(key string, value string) error {
  c, err := getClient()
  if err != nil {
    return err
  }

  result := c.Cmd("set", key, value)
  return result.Err
}

func Load(key string) (string, error) {
  c, err := getClient()
  if err != nil {
    return "", err
  }

  result, err := c.Cmd("get", key).Str()
  return result, err
}
