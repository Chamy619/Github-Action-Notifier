# Github-Action Notifier

Github 액션으로 받은 이벤트 알림

## 사용 방법

### 1. 파일 다운로드 

```sh
wget https://github.com/Chamy619/Github-Action-Notifier/releases/download/v0.0.1/Github-Action-Notifier
```

### 2. 파일 실행 권한 설정

```sh
chmod 755 Github-Action-Notifier
```

### 3. 설정 파일 생성

```yaml
# config.yaml
notifies:
  - files:
      - name: 파일이름
        type: github-action
        added: false
        modified: true
        removed: false
        message: 전송할 메시지
        messenger:
          type: slack
          url: 메시지 전송할 URL
```

- name: 메시지를 받을 타겟 파일의 이름 (경로 포함)
- type: 파일을 조사할 요청 유형 (github-action)
- added: 파일이 추가되었을 때 메시지 수신 여부
- modified: 파일이 변경되었을 때 메시지 수신 여부
- removed: 파일이 제거되었을 때 메시지 수신 여부
- messenger: 메시지를 받을 매체
  - type: 메신저 타입 (slack)
  - url: 메시지 전송할 URL

### 4. 실행

```sh
./Github-Action-Notifier -config=config.yaml -port=8080
```

