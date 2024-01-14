package onprem

import (
    "golang.org/x/crypto/ssh"
    "io/ioutil"
    "net"
    "fmt"
)

func RunCommand(command, hostname, user, keyPath string) (string, error) {
    key, err := ioutil.ReadFile(keyPath)
    if err != nil {
        return "", fmt.Errorf("unable to read private key: %v", err)
    }

    signer, err := ssh.ParsePrivateKey(key)
    if err != nil {
        return "", fmt.Errorf("unable to parse private key: %v", err)
    }

    config := &ssh.ClientConfig{
        User: user,
        Auth: []ssh.AuthMethod{
            ssh.PublicKeys(signer),
        },
        HostKeyCallback: ssh.InsecureIgnoreHostKey(), // Note: In production, use a proper HostKeyCallback
    }

    conn, err := ssh.Dial("tcp", net.JoinHostPort(hostname, "22"), config)
    if err != nil {
        return "", fmt.Errorf("failed to dial: %v", err)
    }
    defer conn.Close()

    session, err := conn.NewSession()
    if err != nil {
        return "", fmt.Errorf("failed to create session: %v", err)
    }
    defer session.Close()

    output, err := session.CombinedOutput(command)
    if err != nil {
        return "", fmt.Errorf("failed to run command: %v", err)
    }

    return string(output), nil
}
