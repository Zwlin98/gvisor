// Copyright 2020 The gVisor Authors.
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

#include "gtest/gtest.h"
#include "test/syscalls/linux/ip_socket_test_util.h"
#include "test/util/capability_util.h"
#include "test/util/test_util.h"
#include "test/util/thread_util.h"

namespace gvisor {
namespace testing {
namespace {

TEST(NetworkNamespaceTest, LoopbackExists) {
  // TODO(b/267210840): Fix this tests for hostinet.
  SKIP_IF(IsRunningWithHostinet());

  SKIP_IF(!ASSERT_NO_ERRNO_AND_VALUE(HaveCapability(CAP_NET_ADMIN)));

  ScopedThread t([&] {
    ASSERT_THAT(unshare(CLONE_NEWNET), SyscallSucceedsWithValue(0));

    // TODO(gvisor.dev/issue/1833): Update this to test that only "lo" exists.
    ASSERT_NE(ASSERT_NO_ERRNO_AND_VALUE(GetLoopbackIndex()), 0);
  });
}

}  // namespace
}  // namespace testing
}  // namespace gvisor
