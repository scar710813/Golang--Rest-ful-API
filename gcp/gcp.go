package gcp

import (
    "cloud.google.com/go/storage"
    "golang.org/x/net/context"
)

func ListBuckets(projectID string) ([]string, error) {
    ctx := context.Background()
    client, err := storage.NewClient(ctx)
    if err != nil {
        return nil, err
    }
    defer client.Close()

    var buckets []string
    it := client.Buckets(ctx, projectID)
    for {
        bkt, err := it.Next()
        if err == iterator.Done {
            break
        }
        if err != nil {
            return nil
