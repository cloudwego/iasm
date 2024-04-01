package aarch64

import (
    `github.com/cloudwego/iasm/asm`
)

const (
    DomainFloat = asm.DomainArchSpecific + iota
    DomainFpSimd
    DomainAdvSimd
    DomainSVE
    DomainSVE2
    DomainSME
    DomainSME2
    DomainSystem
)
