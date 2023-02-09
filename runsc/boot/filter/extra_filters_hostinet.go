// Copyright 2023 The gVisor Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package filter

import (
	"golang.org/x/sys/unix"
	"gvisor.dev/gvisor/pkg/abi/linux"
	"gvisor.dev/gvisor/pkg/seccomp"
)

// hostInetFilters contains syscalls that are needed by sentry/socket/hostinet.
func hostInetFilters() seccomp.SyscallRules {
	return seccomp.SyscallRules{
		unix.SYS_ACCEPT4: []seccomp.Rule{
			{
				seccomp.MatchAny{},
				seccomp.MatchAny{},
				seccomp.MatchAny{},
				seccomp.EqualTo(unix.SOCK_NONBLOCK | unix.SOCK_CLOEXEC),
			},
		},
		unix.SYS_BIND:        {},
		unix.SYS_CONNECT:     {},
		unix.SYS_GETPEERNAME: {},
		unix.SYS_GETSOCKNAME: {},
		unix.SYS_GETSOCKOPT: []seccomp.Rule{
			{
				seccomp.MatchAny{},
				seccomp.EqualTo(unix.SOL_IP),
				seccomp.EqualTo(unix.IP_TOS),
			},
			{
				seccomp.MatchAny{},
				seccomp.EqualTo(unix.SOL_IP),
				seccomp.EqualTo(unix.IP_RECVTOS),
			},
			{
				seccomp.MatchAny{},
				seccomp.EqualTo(unix.SOL_IP),
				seccomp.EqualTo(unix.IP_TTL),
			},
			{
				seccomp.MatchAny{},
				seccomp.EqualTo(unix.SOL_IP),
				seccomp.EqualTo(unix.IP_RECVTTL),
			},
			{
				seccomp.MatchAny{},
				seccomp.EqualTo(unix.SOL_IP),
				seccomp.EqualTo(unix.IP_PKTINFO),
			},
			{
				seccomp.MatchAny{},
				seccomp.EqualTo(unix.SOL_IP),
				seccomp.EqualTo(unix.IP_RECVORIGDSTADDR),
			},
			{
				seccomp.MatchAny{},
				seccomp.EqualTo(unix.SOL_IP),
				seccomp.EqualTo(unix.IP_RECVERR),
			},
			{
				seccomp.MatchAny{},
				seccomp.EqualTo(unix.SOL_IPV6),
				seccomp.EqualTo(unix.IPV6_TCLASS),
			},
			{
				seccomp.MatchAny{},
				seccomp.EqualTo(unix.SOL_IPV6),
				seccomp.EqualTo(unix.IPV6_RECVTCLASS),
			},
			{
				seccomp.MatchAny{},
				seccomp.EqualTo(unix.SOL_IPV6),
				seccomp.EqualTo(unix.IPV6_RECVPKTINFO),
			},
			{
				seccomp.MatchAny{},
				seccomp.EqualTo(unix.SOL_IPV6),
				seccomp.EqualTo(unix.IPV6_UNICAST_HOPS),
			},
			{
				seccomp.MatchAny{},
				seccomp.EqualTo(unix.SOL_IPV6),
				seccomp.EqualTo(unix.IPV6_MULTICAST_HOPS),
			},
			{
				seccomp.MatchAny{},
				seccomp.EqualTo(unix.SOL_IPV6),
				seccomp.EqualTo(unix.IPV6_RECVHOPLIMIT),
			},
			{
				seccomp.MatchAny{},
				seccomp.EqualTo(unix.SOL_IPV6),
				seccomp.EqualTo(unix.IPV6_RECVERR),
			},
			{
				seccomp.MatchAny{},
				seccomp.EqualTo(unix.SOL_IPV6),
				seccomp.EqualTo(unix.IPV6_V6ONLY),
			},
			{
				seccomp.MatchAny{},
				seccomp.EqualTo(unix.SOL_IPV6),
				seccomp.EqualTo(linux.IPV6_RECVORIGDSTADDR),
			},
			{
				seccomp.MatchAny{},
				seccomp.EqualTo(unix.SOL_SOCKET),
				seccomp.EqualTo(unix.SO_BROADCAST),
			},
			{
				seccomp.MatchAny{},
				seccomp.EqualTo(unix.SOL_SOCKET),
				seccomp.EqualTo(unix.SO_ERROR),
			},
			{
				seccomp.MatchAny{},
				seccomp.EqualTo(unix.SOL_SOCKET),
				seccomp.EqualTo(unix.SO_KEEPALIVE),
			},
			{
				seccomp.MatchAny{},
				seccomp.EqualTo(unix.SOL_SOCKET),
				seccomp.EqualTo(unix.SO_SNDBUF),
			},
			{
				seccomp.MatchAny{},
				seccomp.EqualTo(unix.SOL_SOCKET),
				seccomp.EqualTo(unix.SO_RCVBUF),
			},
			{
				seccomp.MatchAny{},
				seccomp.EqualTo(unix.SOL_SOCKET),
				seccomp.EqualTo(unix.SO_REUSEADDR),
			},
			{
				seccomp.MatchAny{},
				seccomp.EqualTo(unix.SOL_SOCKET),
				seccomp.EqualTo(unix.SO_TYPE),
			},
			{
				seccomp.MatchAny{},
				seccomp.EqualTo(unix.SOL_SOCKET),
				seccomp.EqualTo(unix.SO_LINGER),
			},
			{
				seccomp.MatchAny{},
				seccomp.EqualTo(unix.SOL_SOCKET),
				seccomp.EqualTo(unix.SO_TIMESTAMP),
			},
			{
				seccomp.MatchAny{},
				seccomp.EqualTo(unix.SOL_SOCKET),
				seccomp.EqualTo(unix.SO_ACCEPTCONN),
			},
			{
				seccomp.MatchAny{},
				seccomp.EqualTo(unix.SOL_TCP),
				seccomp.EqualTo(unix.TCP_NODELAY),
			},
			{
				seccomp.MatchAny{},
				seccomp.EqualTo(unix.SOL_TCP),
				seccomp.EqualTo(unix.TCP_INFO),
			},
			{
				seccomp.MatchAny{},
				seccomp.EqualTo(unix.SOL_TCP),
				seccomp.EqualTo(linux.TCP_INQ),
			},
			{
				seccomp.MatchAny{},
				seccomp.EqualTo(unix.SOL_TCP),
				seccomp.EqualTo(linux.TCP_MAXSEG),
			},
			{
				seccomp.MatchAny{},
				seccomp.EqualTo(unix.SOL_TCP),
				seccomp.EqualTo(linux.TCP_CONGESTION),
			},
			{
				seccomp.MatchAny{},
				seccomp.EqualTo(unix.SOL_TCP),
				seccomp.EqualTo(linux.TCP_USER_TIMEOUT),
			},
			{
				seccomp.MatchAny{},
				seccomp.EqualTo(unix.SOL_TCP),
				seccomp.EqualTo(linux.TCP_DEFER_ACCEPT),
			},
			{
				seccomp.MatchAny{},
				seccomp.EqualTo(unix.SOL_TCP),
				seccomp.EqualTo(linux.TCP_SYNCNT),
			},
			{
				seccomp.MatchAny{},
				seccomp.EqualTo(unix.SOL_TCP),
				seccomp.EqualTo(linux.TCP_WINDOW_CLAMP),
			},
		},
		unix.SYS_IOCTL: []seccomp.Rule{
			{
				seccomp.MatchAny{},
				seccomp.EqualTo(unix.TIOCOUTQ),
			},
			{
				seccomp.MatchAny{},
				seccomp.EqualTo(unix.TIOCINQ),
			},
			{
				seccomp.MatchAny{},
				seccomp.EqualTo(unix.SIOCGIFFLAGS),
			},
			{
				seccomp.MatchAny{},
				seccomp.EqualTo(unix.SIOCGIFCONF),
			},
		},
		unix.SYS_LISTEN:   {},
		unix.SYS_READV:    {},
		unix.SYS_RECVFROM: {},
		unix.SYS_RECVMSG:  {},
		unix.SYS_SENDMSG:  {},
		unix.SYS_SENDTO:   {},
		unix.SYS_SETSOCKOPT: []seccomp.Rule{
			{
				seccomp.MatchAny{},
				seccomp.EqualTo(unix.SOL_SOCKET),
				seccomp.EqualTo(unix.SO_BROADCAST),
				seccomp.MatchAny{},
				seccomp.EqualTo(4),
			},
			{
				seccomp.MatchAny{},
				seccomp.EqualTo(unix.SOL_SOCKET),
				seccomp.EqualTo(unix.SO_SNDBUF),
				seccomp.MatchAny{},
				seccomp.EqualTo(4),
			},
			{
				seccomp.MatchAny{},
				seccomp.EqualTo(unix.SOL_SOCKET),
				seccomp.EqualTo(unix.SO_RCVBUF),
				seccomp.MatchAny{},
				seccomp.EqualTo(4),
			},
			{
				seccomp.MatchAny{},
				seccomp.EqualTo(unix.SOL_SOCKET),
				seccomp.EqualTo(unix.SO_REUSEADDR),
				seccomp.MatchAny{},
				seccomp.EqualTo(4),
			},
			{
				seccomp.MatchAny{},
				seccomp.EqualTo(unix.SOL_SOCKET),
				seccomp.EqualTo(unix.SO_TIMESTAMP),
				seccomp.MatchAny{},
				seccomp.EqualTo(4),
			},
			{
				seccomp.MatchAny{},
				seccomp.EqualTo(unix.SOL_TCP),
				seccomp.EqualTo(unix.TCP_NODELAY),
				seccomp.MatchAny{},
				seccomp.EqualTo(4),
			},
			{
				seccomp.MatchAny{},
				seccomp.EqualTo(unix.SOL_TCP),
				seccomp.EqualTo(linux.TCP_INQ),
				seccomp.MatchAny{},
				seccomp.EqualTo(4),
			},
			{
				seccomp.MatchAny{},
				seccomp.EqualTo(unix.SOL_TCP),
				seccomp.EqualTo(linux.TCP_MAXSEG),
				seccomp.MatchAny{},
				seccomp.EqualTo(4),
			},
			{
				seccomp.MatchAny{},
				seccomp.EqualTo(unix.SOL_TCP),
				seccomp.EqualTo(linux.TCP_CONGESTION),
			},
			{
				seccomp.MatchAny{},
				seccomp.EqualTo(unix.SOL_TCP),
				seccomp.EqualTo(linux.TCP_USER_TIMEOUT),
			},
			{
				seccomp.MatchAny{},
				seccomp.EqualTo(unix.SOL_TCP),
				seccomp.EqualTo(linux.TCP_DEFER_ACCEPT),
			},
			{
				seccomp.MatchAny{},
				seccomp.EqualTo(unix.SOL_TCP),
				seccomp.EqualTo(linux.TCP_SYNCNT),
			},
			{
				seccomp.MatchAny{},
				seccomp.EqualTo(unix.SOL_TCP),
				seccomp.EqualTo(linux.TCP_WINDOW_CLAMP),
			},
			{
				seccomp.MatchAny{},
				seccomp.EqualTo(unix.SOL_IP),
				seccomp.EqualTo(unix.IP_TOS),
				seccomp.MatchAny{},
				seccomp.EqualTo(4),
			},
			{
				seccomp.MatchAny{},
				seccomp.EqualTo(unix.SOL_IP),
				seccomp.EqualTo(unix.IP_RECVTOS),
				seccomp.MatchAny{},
				seccomp.EqualTo(4),
			},
			{
				seccomp.MatchAny{},
				seccomp.EqualTo(unix.SOL_IP),
				seccomp.EqualTo(unix.IP_TTL),
				seccomp.MatchAny{},
				seccomp.EqualTo(4),
			},
			{
				seccomp.MatchAny{},
				seccomp.EqualTo(unix.SOL_IP),
				seccomp.EqualTo(unix.IP_RECVTTL),
				seccomp.MatchAny{},
				seccomp.EqualTo(4),
			},
			{
				seccomp.MatchAny{},
				seccomp.EqualTo(unix.SOL_IP),
				seccomp.EqualTo(unix.IP_PKTINFO),
				seccomp.MatchAny{},
				seccomp.EqualTo(4),
			},
			{
				seccomp.MatchAny{},
				seccomp.EqualTo(unix.SOL_IPV6),
				seccomp.EqualTo(unix.IPV6_RECVPKTINFO),
				seccomp.MatchAny{},
				seccomp.EqualTo(4),
			},
			{
				seccomp.MatchAny{},
				seccomp.EqualTo(unix.SOL_IP),
				seccomp.EqualTo(unix.IP_RECVORIGDSTADDR),
				seccomp.MatchAny{},
				seccomp.EqualTo(4),
			},
			{
				seccomp.MatchAny{},
				seccomp.EqualTo(unix.SOL_IP),
				seccomp.EqualTo(unix.IP_RECVERR),
				seccomp.MatchAny{},
				seccomp.EqualTo(4),
			},
			{
				seccomp.MatchAny{},
				seccomp.EqualTo(unix.SOL_IPV6),
				seccomp.EqualTo(unix.IPV6_TCLASS),
				seccomp.MatchAny{},
				seccomp.EqualTo(4),
			},
			{
				seccomp.MatchAny{},
				seccomp.EqualTo(unix.SOL_IPV6),
				seccomp.EqualTo(unix.IPV6_RECVTCLASS),
				seccomp.MatchAny{},
				seccomp.EqualTo(4),
			},
			{
				seccomp.MatchAny{},
				seccomp.EqualTo(unix.SOL_IPV6),
				seccomp.EqualTo(unix.IPV6_UNICAST_HOPS),
				seccomp.MatchAny{},
				seccomp.EqualTo(4),
			},
			{
				seccomp.MatchAny{},
				seccomp.EqualTo(unix.SOL_IPV6),
				seccomp.EqualTo(unix.IPV6_MULTICAST_HOPS),
				seccomp.MatchAny{},
				seccomp.EqualTo(4),
			},
			{
				seccomp.MatchAny{},
				seccomp.EqualTo(unix.SOL_IPV6),
				seccomp.EqualTo(unix.IPV6_RECVHOPLIMIT),
				seccomp.MatchAny{},
				seccomp.EqualTo(4),
			},
			{
				seccomp.MatchAny{},
				seccomp.EqualTo(unix.SOL_IPV6),
				seccomp.EqualTo(linux.IPV6_RECVORIGDSTADDR),
				seccomp.MatchAny{},
				seccomp.EqualTo(4),
			},
			{
				seccomp.MatchAny{},
				seccomp.EqualTo(unix.SOL_IPV6),
				seccomp.EqualTo(unix.IPV6_RECVERR),
				seccomp.MatchAny{},
				seccomp.EqualTo(4),
			},
			{
				seccomp.MatchAny{},
				seccomp.EqualTo(unix.SOL_IPV6),
				seccomp.EqualTo(unix.IPV6_V6ONLY),
				seccomp.MatchAny{},
				seccomp.EqualTo(4),
			},
		},
		unix.SYS_SHUTDOWN: []seccomp.Rule{
			{
				seccomp.MatchAny{},
				seccomp.EqualTo(unix.SHUT_RD),
			},
			{
				seccomp.MatchAny{},
				seccomp.EqualTo(unix.SHUT_WR),
			},
			{
				seccomp.MatchAny{},
				seccomp.EqualTo(unix.SHUT_RDWR),
			},
		},
		unix.SYS_SOCKET: []seccomp.Rule{
			{
				seccomp.EqualTo(unix.AF_INET),
				seccomp.EqualTo(unix.SOCK_STREAM | unix.SOCK_NONBLOCK | unix.SOCK_CLOEXEC),
				seccomp.EqualTo(0),
			},
			{
				seccomp.EqualTo(unix.AF_INET),
				seccomp.EqualTo(unix.SOCK_DGRAM | unix.SOCK_NONBLOCK | unix.SOCK_CLOEXEC),
				seccomp.EqualTo(0),
			},
			{
				seccomp.EqualTo(unix.AF_INET6),
				seccomp.EqualTo(unix.SOCK_STREAM | unix.SOCK_NONBLOCK | unix.SOCK_CLOEXEC),
				seccomp.EqualTo(0),
			},
			{
				seccomp.EqualTo(unix.AF_INET6),
				seccomp.EqualTo(unix.SOCK_DGRAM | unix.SOCK_NONBLOCK | unix.SOCK_CLOEXEC),
				seccomp.EqualTo(0),
			},
		},
		unix.SYS_WRITEV: {},
	}
}
