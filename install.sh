#/bin/bash
# The script requires root permissions

function install_trento() {
    . /etc/os-release
    if [[ ! $NAME =~ "SUSE" ]]; then
        echo "Not SUSE operating system."
        return -1
    fi
    case "$VERSION_ID" in
    "15.3")
        repo="https://download.opensuse.org/repositories/devel:/sap:/trento/15.3/devel:sap:trento.repo"
        ;;
    *)
        echo "Not supported OS release."
        return -1;;
    esac
    path=${repo%/*}/
    if zypper lr --details | cut -d'|' -f9 | grep $path; then
        echo "$path repository alreadt exists. Skipping."
    else
        zypper ar $repo
    fi
    zypper ref
    if which trento; then
        echo "trento is already installed. Skipping."
    else
        zypper in trento
    fi
}

function install_consul() {
    if which consul; then
        echo "consul is already installed. Skipping."
        return 0
    fi
    # Unzip cannot be extracted of the fly, so we store the file, extract and remove.
    zipfile=$(mktemp)
    # FIXME! The version is hardcoded. It should take the last existing.
    wget https://releases.hashicorp.com/consul/1.9.8/consul_1.9.8_linux_amd64.zip -O $zipfile
    unzip $zipfile -d /usr/bin/
    rm $zipfile
}

function install_all() {
    install_trento
    rc=$?
    if (( $rc != 0 )); then
        return $rc
    fi
    install_consul
    rc=$?
    if (( $rc != 0 )); then
        return $rc
    fi
    echo "Installation successfully completed."
    return 0
}

function print_help() {
    cat <<END
This is an installer of trento. Trento is a web-based graphical user interface
that can help you deploy, provision and operate infrastructure for SAP Applications

Usage:

  ./install.sh [command]

Available Commands:
  all         Installs trento and consul.
  trento      Installs trento.
  consul      Installs consul.
  help        Prints this help information
  version     Prints the version of Trento
END
}

function print_version() {
    echo "Trento version 0"
}

ACTION=$1

# These operations don't require OCF parameters to be set
case "$ACTION" in
    all)     install_all
             exit $?;;
    trento)  install_trento
             exit $?;;
    consul)  install_consul
             exit $?;;
    help)    print_help
             exit $?;;
    version) print_version
             exit $?;;
    *)       install_all
             exit $?;;
esac
