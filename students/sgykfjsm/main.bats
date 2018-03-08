#!/usr/bin/env bats

task_cmd="bin/task"

echo_error(){
    echo -e "ERROR: $*"
}

@test "task add should be worked" {
    tasks=("review talk proposal" "clean dishes")
    for task in "${tasks[@]}"
    do
        run ${task_cmd} add ${task}
        if [ "${status}" -ne 0 ]; then
            echo_error "failed to add task. status is ${status}"
            return 1
        fi

        expected="Added \"${task}\" to your task list."
        if [ "${output}" != "${expected}" ]; then
            echo_error "expected is '${expected}', but actual is '${output}'."
            return 1
        fi
    done
}

@test "task list should show the unfinished tasks (1)" {
    run ${task_cmd} list
    if [ "${status}" -ne 0 ]; then
        echo_error "failed to show tasks. status is ${status}"
        return 1
    fi

    expected="You have the following tasks:
1. review talk proposal
2. clean dishes"
    if [ "${output}" != "${expected}" ]; then
        echo_error "expected is '${expected}', but actual is '${output}'."
        return 1
    fi
}

@test "task do should mark task as finished with given task number" {
    run ${task_cmd} do 1
    if [ "${status}" -ne 0 ]; then
        echo_error "failed to show tasks. status is ${status}"
        return 1
    fi

    expected='You have completed the "review talk proposal" task.'
    if [ "${output}" != "${expected}" ]; then
        echo_error "expected is '${expected}', but actual is '${output}'."
        return 1
    fi
}

@test "task list should show the unfinished tasks (2)" {
    run ${task_cmd} list
    if [ "${status}" -ne 0 ]; then
        echo_error "failed to show tasks. status is ${status}"
        return 1
    fi

    expected="You have the following tasks:
1. clean dishes"
    if [ "${output}" != "${expected}" ]; then
        echo_error "expected is '${expected}', but actual is '${output}'."
        return 1
    fi
}
